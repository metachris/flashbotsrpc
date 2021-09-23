package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.

func main() {
	rpc := flashbotsrpc.NewFlashbotsRPC("https://relay.flashbots.net")

	// Query relay for user stats
	result, err := rpc.FlashbotsGetUserStats(privateKey, 13281018)
	if err != nil {
		panic(err)
	}

	// Print result
	fmt.Printf("%+v\n", result)
}
