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
	urls := []string{
		"https://relay.flashbots.net",
		"https://rpc.titanbuilder.xyz",
		"https://builder0x69.io",
		"https://rpc.beaverbuild.org",
		"https://rsync-builder.xyz",
		"https://api.blocknative.com/v1/auction",
		// "https://mev.api.blxrbdn.com", # Authentication required
		"https://eth-builder.com",
		"https://builder.gmbit.co/rpc",
		"https://buildai.net",
		"https://rpc.payload.de",
		"https://rpc.lightspeedbuilder.info",
		"https://rpc.nfactorial.xyz",
	}

	rpc := flashbotsrpc.NewBuilderBroadcastRPC(urls)
	rpc.Debug = true

	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{"YOUR_RAW_TX"},
		BlockNumber: fmt.Sprintf("0x%x", 13281018),
	}

	results := rpc.BroadcastBundle(privateKey, sendBundleArgs)
	for _, result := range results {
		if result.Err != nil {
			if errors.Is(result.Err, flashbotsrpc.ErrRelayErrorResponse) {
				// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
				fmt.Println(result.Err.Error())
			} else {
				fmt.Printf("error: %+v\n", result.Err)
			}
			return
		}

		// Print result
		fmt.Printf("%+v\n", result.BundleResponse)
	}
}
