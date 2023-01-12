package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"code.vegaprotocol.io/tranches/contracts"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/maps"
)

const (
	deploymentHeight  uint64 = 12834524
	catchupBlocksSize uint64 = 50000
)

// Events
const (
	trancheCreated        = "Tranche_Created"
	trancheBalanceAdded   = "Tranche_Balance_Added"
	trancheBalanceRemoved = "Tranche_Balance_Removed"
)

var (
	port        uint
	ethereumRPC string

	vestingContractAddress = common.HexToAddress("0x23d1bFE8fA50a167816fBD79D7932577c06011f4")
	pollEventRetryDuration = 14 * time.Second // pretty much how long it takes to build a block
	pollNewState           = time.Minute      // no need to poll too often
)

func init() {
	flag.UintVar(&port, "port", 1789, "port of the http server")
	flag.StringVar(&ethereumRPC, "eth-rpc", "", "ethereum RPC address")
}

type Tranche struct {
	InitialBalance *big.Int                 `json:"initial_balance"`
	CurrentBalance *big.Int                 `json:"current_balance"`
	Users          map[common.Address]*User `json:"users"`
	CliffStart     *big.Int                 `json:"cliff_start"`
	Duration       *big.Int                 `json:"duration"`
}

type User struct {
	TranchID    uint8                `json:"tranch_id"`
	Balance     *big.Int             `json:"balance"`
	RemovedLogs []BalanceMovementLog `json:"removed_logs"`
	AddedLogs   []BalanceMovementLog `json:"added_logs"`
}

func (u *User) IntoResponse() *UserResponse {
	removedLogs := make([]BalanceMovementLogResponse, 0, len(u.RemovedLogs))
	for _, v := range u.RemovedLogs {
		removedLogs = append(removedLogs, v.IntoResponse())
	}

	addedLogs := make([]BalanceMovementLogResponse, 0, len(u.AddedLogs))
	for _, v := range u.AddedLogs {
		addedLogs = append(addedLogs, v.IntoResponse())
	}

	return &UserResponse{
		TranchID:    u.TranchID,
		Balance:     u.Balance.String(),
		AddedLogs:   addedLogs,
		RemovedLogs: removedLogs,
	}
}

type BalanceMovementLog struct {
	TrancheID   uint8          `json:"tranche_id"`
	Address     common.Address `json:"address"`
	TxHash      string         `json:"tx_hash"`
	Amount      *big.Int       `json:"amount"`
	BlockNumber uint64         `json:"block_number"`
}

func (b *BalanceMovementLog) IntoResponse() BalanceMovementLogResponse {
	return BalanceMovementLogResponse{
		TrancheID:   b.TrancheID,
		Address:     b.Address,
		TxHash:      b.TxHash,
		Amount:      b.Amount.String(),
		BlockNumber: b.BlockNumber,
	}
}

type State struct {
	clt      *ethclient.Client
	filterer *contracts.ERC20VestingFilterer
	abi      ethabi.ABI

	currentHeightLastUpdate time.Time
	currentHeight           uint64
	lastHeightPulled        uint64

	// tanche => ( address => user)
	mu    sync.RWMutex
	store map[uint8]*Tranche
}

func NewState(clt *ethclient.Client) *State {
	filterer, err := contracts.NewERC20VestingFilterer(vestingContractAddress, clt)
	if err != nil {
		log.Fatalf("couldn't create log filterer for ERC20 brigde: %v", err)
	}

	abi, err := ethabi.JSON(strings.NewReader(contracts.ERC20VestingMetaData.ABI))
	if err != nil {
		log.Fatalf("couldn't load collateral bridge ABI: %v", err)
	}

	tranches := map[uint8]*Tranche{}
	// these seems to exists without creating them
	for _, v := range []uint8{0, 154, 110, 107, 66} {
		tranches[v] = &Tranche{
			InitialBalance: big.NewInt(0),
			CurrentBalance: big.NewInt(0),
			CliffStart:     big.NewInt(0),
			Duration:       big.NewInt(0),
			Users:          map[common.Address]*User{},
		}
	}

	return &State{
		clt:      clt,
		filterer: filterer,
		abi:      abi,
		store:    tranches,
	}
}

func (s *State) CurrentHeight(ctx context.Context) uint64 {
	currentHeight := new(uint64)
	*currentHeight = s.currentHeight

	if now := time.Now(); s.currentHeightLastUpdate.Add(15).Before(now) {
		infiniteRetry(func() error {
			lastBlockHeader, err := s.clt.HeaderByNumber(ctx, nil)
			if err != nil {
				return fmt.Errorf("couldn't get the current height of Ethereum blockchain: %v", err)
			}

			s.currentHeightLastUpdate = now
			s.currentHeight = lastBlockHeader.Number.Uint64()
			*currentHeight = s.currentHeight

			return nil
		}, pollEventRetryDuration)
	}

	return *currentHeight
}

func (s *State) newQuery(startAt uint64, stopAt uint64) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startAt),
		ToBlock:   new(big.Int).SetUint64(stopAt),
		Addresses: []common.Address{
			vestingContractAddress,
		},
		Topics: [][]common.Hash{{
			s.abi.Events[trancheCreated].ID,
			s.abi.Events[trancheBalanceAdded].ID,
			s.abi.Events[trancheBalanceRemoved].ID,
		}},
	}

	return query
}

func (s *State) filterLogs(ctx context.Context, query ethereum.FilterQuery) []ethtypes.Log {
	var logs []ethtypes.Log

	infiniteRetry(func() error {
		l, err := s.clt.FilterLogs(ctx, query)
		if err != nil {
			log.Printf("error: couldn't subscribe to the Ethereum log filterer: %v", err)
			return fmt.Errorf("couldn't subscribe to the Ethereum log filterer: %w", err)
		}
		logs = l
		return nil
	}, pollEventRetryDuration)

	return logs
}

func (s *State) createTranche(event *contracts.ERC20VestingTrancheCreated) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[event.TrancheId] = &Tranche{
		InitialBalance: big.NewInt(0),
		CurrentBalance: big.NewInt(0),
		CliffStart:     new(big.Int).Set(event.CliffStart),
		Duration:       new(big.Int).Set(event.Duration),
		Users:          map[common.Address]*User{},
	}
}

func (s *State) addBalance(event *contracts.ERC20VestingTrancheBalanceAdded) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tranche, ok := s.store[event.TrancheId]
	if !ok {
		log.Fatalf("unknow tranche id in balance added event: %v", event.TrancheId)
	}

	user, ok := tranche.Users[event.User]
	if !ok {
		user = &User{
			TranchID:    event.TrancheId,
			Balance:     big.NewInt(0),
			RemovedLogs: []BalanceMovementLog{},
			AddedLogs:   []BalanceMovementLog{},
		}
	}

	user.Balance.Add(user.Balance, event.Amount)
	user.AddedLogs = append(user.AddedLogs, BalanceMovementLog{
		TrancheID:   event.TrancheId,
		Address:     event.User,
		Amount:      new(big.Int).Set(event.Amount),
		TxHash:      event.Raw.TxHash.Hex(),
		BlockNumber: event.Raw.BlockNumber,
	})

	tranche.Users[event.User] = user

	tranche.CurrentBalance.Add(tranche.CurrentBalance, event.Amount)
	tranche.InitialBalance.Add(tranche.InitialBalance, event.Amount)
	s.store[event.TrancheId] = tranche
}

func (s *State) removeBalance(event *contracts.ERC20VestingTrancheBalanceRemoved) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tranche, ok := s.store[event.TrancheId]
	if !ok {
		log.Fatalf("unknow tranche id in balance removed event: %v", event.TrancheId)
	}

	user, ok := tranche.Users[event.User]
	if !ok {
		user = &User{
			TranchID:    event.TrancheId,
			Balance:     big.NewInt(0),
			RemovedLogs: []BalanceMovementLog{},
			AddedLogs:   []BalanceMovementLog{},
		}
	}

	user.Balance.Sub(user.Balance, event.Amount)
	user.RemovedLogs = append(user.RemovedLogs, BalanceMovementLog{
		TrancheID:   event.TrancheId,
		Address:     event.User,
		Amount:      new(big.Int).Set(event.Amount),
		TxHash:      event.Raw.TxHash.Hex(),
		BlockNumber: event.Raw.BlockNumber,
	})

	tranche.Users[event.User] = user

	tranche.CurrentBalance.Sub(tranche.CurrentBalance, event.Amount)
	s.store[event.TrancheId] = tranche
}

func (s *State) parseEvent(l ethtypes.Log) {

	switch l.Topics[0] {
	case s.abi.Events[trancheCreated].ID:
		event, err := s.filterer.ParseTrancheCreated(l)
		if err != nil {
			log.Fatalf("Couldn't parse AssetDeposited event: %v", err)
		}
		s.createTranche(event)
	case s.abi.Events[trancheBalanceAdded].ID:
		event, err := s.filterer.ParseTrancheBalanceAdded(l)
		if err != nil {
			log.Fatalf("Couldn't parse AssetDeposited event: %v", err)
		}
		s.addBalance(event)
	case s.abi.Events[trancheBalanceRemoved].ID:
		event, err := s.filterer.ParseTrancheBalanceRemoved(l)
		if err != nil {
			log.Fatalf("Couldn't parse AssetDeposited event: %v", err)
		}
		s.removeBalance(event)
	}
}

func (s *State) CatchupSinceDeployment() {

	var (
		from   = deploymentHeight
		to     = deploymentHeight
		height = s.CurrentHeight(context.Background())
	)

	log.Printf("catching up state from contract deployment(%v) to current ethereum height(%v)", deploymentHeight, height)

	for {
		from, to = to, to+catchupBlocksSize

		if to > height {
			to = height
		}

		log.Printf("loading from block %v to %v", from, to)

		query := s.newQuery(from, to)
		logs := s.filterLogs(context.Background(), query)

		for _, l := range logs {
			s.parseEvent(l)
		}

		if to == height {
			s.lastHeightPulled = to
			// we are done
			break
		}
		// increase to start from next block then
		to += 1
	}
}

func (s *State) PullForever() {
	ticker := time.NewTicker(pollNewState)
	for range ticker.C {
		if newHeight := s.CurrentHeight(context.Background()); newHeight > s.lastHeightPulled {
			from, to := s.lastHeightPulled+1, newHeight
			log.Printf("loading from block %v to %v", from, to)

			query := s.newQuery(from, to)
			logs := s.filterLogs(context.Background(), query)

			for _, l := range logs {
				s.parseEvent(l)
			}

			s.lastHeightPulled = to
		}
	}
}

type UserByAddressResponse struct {
	Tranches map[uint8]*UserResponse `json:"tranches"`
}

type UserResponse struct {
	TranchID    uint8                        `json:"tranch_id"`
	Balance     string                       `json:"balance"`
	RemovedLogs []BalanceMovementLogResponse `json:"removed_logs"`
	AddedLogs   []BalanceMovementLogResponse `json:"added_logs"`
}

type BalanceMovementLogResponse struct {
	TrancheID   uint8          `json:"tranche_id"`
	Address     common.Address `json:"address"`
	TxHash      string         `json:"tx_hash"`
	Amount      string         `json:"amount"`
	BlockNumber uint64         `json:"block_number"`
}

func (s *State) GetUserByAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	address := ps.ByName("address")
	if !common.IsHexAddress(address) {
		writeError(w, http.StatusBadRequest, errors.New("not a valid ethereum address"))
		return
	}

	ethAddress := common.HexToAddress(address)
	tranches := map[uint8]*UserResponse{}

	for k, v := range s.store {
		user, ok := v.Users[ethAddress]
		if !ok {
			// this user is not in this tranche, move on
			continue
		}
		tranches[k] = user.IntoResponse()
	}

	writeOK(w, UserByAddressResponse{Tranches: tranches})
}

type TrancheStats struct {
	TrancheID      uint8            `json:"tranche_id"`
	Users          []common.Address `json:"users"`
	InitialBalance *big.Int         `json:"initial_balance"`
	CurrentBalance *big.Int         `json:"current_balance"`
	CliffStart     *big.Int         `json:"cliff_start"`
	Duration       *big.Int         `json:"duration"`
}

type TranchesStatsResponse struct {
	Tranches map[uint8]TrancheStats `json:"tranches"`
}

func (s *State) GetTranchesStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tranches := map[uint8]TrancheStats{}
	for k, v := range s.store {
		tranches[k] = TrancheStats{
			TrancheID:      k,
			Users:          maps.Keys(v.Users),
			InitialBalance: v.InitialBalance,
			CurrentBalance: v.CurrentBalance,
			CliffStart:     v.CliffStart,
			Duration:       v.Duration,
		}
	}

	writeOK(w, TranchesStatsResponse{Tranches: tranches})
}

func main() {
	flag.Parse()

	if len(ethereumRPC) <= 0 {
		log.Fatal("error: missing eth-rpc parameter")
	}

	ethClient, err := ethclient.DialContext(context.Background(), ethereumRPC)
	if err != nil {
		log.Fatalf("couldn't instantiate Ethereum client: %v", err)
	}

	state := NewState(ethClient)

	log.Print("loading current state...")
	state.CatchupSinceDeployment()
	log.Print("current state loaded successfully")

	log.Printf("starting http API on port: %v", port)

	router := httprouter.New()
	router.GET("/user/:address", state.GetUserByAddress)
	router.GET("/tranches/stats", state.GetTranchesStats)

	go state.PullForever()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}

func writeOK(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	buf, err := json.Marshal(payload)
	if err != nil {
		log.Printf("couldn't marshal response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(buf); err != nil {
		log.Printf("couldn't write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	buf, err := json.Marshal(&struct {
		Error string `json:"error"`
	}{Error: err.Error()})
	if err != nil {
		log.Printf("couldn't marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(buf); err != nil {
		log.Printf("couldn't write error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// We are retrying infinitely, on purpose, as we don't want the Ethereum
// Forwarder to exit, and this under any circumstances. Failure is not an option.
func infiniteRetry(fn backoff.Operation, durationBetweenTwoRetry time.Duration) {
	// No need to retrieve the error, as we are waiting indefinitely for a
	// success.
	_ = backoff.Retry(fn, backoff.NewConstantBackOff(durationBetweenTwoRetry))
}
