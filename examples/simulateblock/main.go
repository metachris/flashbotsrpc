package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/flashbotsrpc"
)

func main() {
	mevGethUriPtr := flag.String("mevgeth", os.Getenv("MEVGETH_NODE"), "mev-geth node URI")
	blockHash := flag.String("blockhash", "", "hash of block to simulate")
	blockNumber := flag.Int64("blocknumber", -1, "number of block to simulate")
	debugPtr := flag.Bool("debug", false, "print debug information")
	flag.Parse()

	if *mevGethUriPtr == "" {
		log.Fatal("No mev geth URI provided")
	}

	if *blockHash == "" && *blockNumber == -1 {
		log.Fatal("Either block number or hash is needed")
	}

	client, err := ethclient.Dial(*mevGethUriPtr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to", *mevGethUriPtr)

	var block *types.Block
	if *blockHash != "" {
		hash := common.HexToHash(*blockHash)
		block, err = client.BlockByHash(context.Background(), hash)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		block, err = client.BlockByNumber(context.Background(), big.NewInt(*blockNumber))
		if err != nil {
			log.Fatal(err)
		}
	}

	t := time.Unix(int64(block.Header().Time), 0).UTC()
	fmt.Printf("Block %d %s \t %s \t tx=%-4d \t gas=%d \t uncles=%d\n", block.NumberU64(), block.Hash(), t, len(block.Transactions()), block.GasUsed(), len(block.Uncles()))

	if len(block.Transactions()) == 0 {
		fmt.Println("No transactions in this block")
		return
	}

	rpc := flashbotsrpc.New(*mevGethUriPtr)
	rpc.Debug = *debugPtr

	privateKey, _ := crypto.GenerateKey()
	result, err := rpc.FlashbotsSimulateBlock(privateKey, block, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CoinbaseDiff:", result.CoinbaseDiff)
}
