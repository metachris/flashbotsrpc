package main

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.
// var privateKey, _ = crypto.HexToECDSA("YOUR_PRIVATE_KEY")

func main() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	cancelPrivTxArgs := flashbotsrpc.FlashbotsCancelPrivateTransactionRequest{
		TxHash: "0xYOUR_TX_HASH",
	}

	cancelled, err := rpc.FlashbotsCancelPrivateTransaction(privateKey, cancelPrivTxArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	// Print result
	fmt.Printf("was cancelled: %v\n", cancelled)
}
