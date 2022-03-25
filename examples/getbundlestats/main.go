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

	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{"YOUR_RAW_TX"},
		BlockNumber: fmt.Sprintf("0x%x", 13281018),
	}

	result, err := rpc.FlashbotsSendBundle(privateKey, sendBundleArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	getBundleStatsArgs := flashbotsrpc.FlashbotsGetBundleStatsParam{
		BlockNumber: fmt.Sprintf("0x%x", 13281018),
		BundleHash:  result.BundleHash,
	}
	bundleStats, err := rpc.FlashbotsGetBundleStats(privateKey, getBundleStatsArgs)
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
	fmt.Printf("%+v\n", bundleStats)
}
