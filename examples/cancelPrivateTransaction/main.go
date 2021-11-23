package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.

func main() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	cancelPrivTxArgs := flashbotsrpc.FlashbotsCancelPrivateTransactionRequest{
		TxHash: "0xfb34b88cd77215867aa8e8ff0abc7060178b8fed6519a85d0b22853dfd5e9fec",
	}

	cancelled, err := rpc.FlashbotsCancelPrivateTransaction(privateKey, cancelPrivTxArgs)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// Print result
	fmt.Printf("was cancelled: %v\n", cancelled)
}
