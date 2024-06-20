# Tranches

## How to run it:

1. Build it:

```shell
go build -o tranches ./main.go
```

2. Run it 

```shell
./tranches
```

## Flags

- `--port` - port of the http server
- `--eth-rpc` - ethereum RPC address
- `--vesting-contract-address` - vesting contract address
- `--deployment-height` - the deployment block height for the given contract

	flag.UintVar(&port, "port", 1789, "port of the http server")
	flag.StringVar(&ethereumRPC, "eth-rpc", "", "ethereum RPC address")
	flag.StringVar(&vestingContractAddress, "vesting-contract-address", "0x23d1bFE8fA50a167816fBD79D7932577c06011f4", "vesting contract address")
	flag.Uint64Var(&deploymentHeight, "deployment-height", 12834524, "deployment block height")
}