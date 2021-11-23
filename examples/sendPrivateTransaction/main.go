package main

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.

func main() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	sendPrivTxArgs := flashbotsrpc.FlashbotsSendPrivateTransactionRequest{
		Tx: "0xYourTxHash",
	}

	txHash, err := rpc.FlashbotsSendPrivateTransaction(privateKey, sendPrivTxArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			fmt.Println(err.Error()) // standard error response from relay
		} else {
			fmt.Printf("unknown error: %+v\n", err)
		}
		return
	}

	// Print txHash
	fmt.Printf("txHash: %s\n", txHash)
}
