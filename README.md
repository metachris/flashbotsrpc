# Flashbots RPC client

Fork of [ethrpc](https://github.com/onrik/ethrpc) with additional [Flashbots RPC methods](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint):

* `FlashbotsCallBundle` ([`eth_callBundle`](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint/#eth_callbundle))
* `FlashbotsSendBundle` ([`eth_sendBundle`](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint/#eth_sendbundle))
* `FlashbotsGetUserStats` ([`flashbots_getUserStats`](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint/#flashbots_getuserstats))
* `FlashbotsSendPrivateTransaction` (`eth_sendPrivateTransaction`)
* `FlashbotsCancelPrivateTransaction` (`eth_cancelPrivateTransaction`)
* `FlashbotsSimulateBlock`: simulate a full block

## Usage

`go get github.com/metachris/flashbotsrpc`

```go
rpc := flashbotsrpc.New("https://relay.flashbots.net")

// Creating a new private key here for testing, you probably want to use an existing one
privateKey, _ := crypto.GenerateKey()

// flashbots_getUserStats example
// ------------------------------
result, err := rpc.FlashbotsGetUserStats(privateKey, 13281018)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", result)

// eth_sendBundle example
// ----------------------
sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
    Txs:         []string{"YOUR_HASH"},
    BlockNumber: fmt.Sprintf("0x%x", 13281018),
}

result, err := rpc.FlashbotsSendBundle(privateKey, sendBundleArgs)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", result)
```

You can find example code in the [`/examples/` directory](https://github.com/metachris/flashbotsrpc/tree/master/examples).
