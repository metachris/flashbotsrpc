# Flashbots RPC client

Fork of [ethrpc](https://github.com/onrik/ethrpc) with additional [Flashbots RPC methods](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint):

* `FlashbotsGetUserStats`
* `FlashbotsCallBundle`
* `FlashbotsSendBundle`
* `FlashbotsSimulateBlock`: simulate a full block

## Usage

```go
rpc := flashbotsrpc.New("https://relay.flashbots.net")

// Creating a new private key here for testing; you probably would want to use an existing one
privateKey, _ := crypto.GenerateKey() 

// Query relay for user stats
result, err := rpc.FlashbotsGetUserStats(privateKey, 13281018)
if err != nil {
    log.Fatal(err)
}

// Print result
fmt.Printf("%+v\n", result)
```

You can find [more examples in `/examples/`](https://github.com/metachris/flashbotsrpc/tree/master/examples).
