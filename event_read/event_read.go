package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("wss://arb-sepolia.g.alchemy.com/v2/GBDdpIse13KF_izTI2-J5KACwNNtxS0Z")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x7C44AfdCcc6857338bd30CE9CD0C6c3788309777")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(77146703),
		ToBlock:   big.NewInt(77825595),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("BlockHash ", vLog.BlockHash.Hex())
		fmt.Println("BlockNumber ", vLog.BlockNumber)
		fmt.Println("TxHash ", vLog.TxHash.Hex())

		a := new(big.Int)
		a.SetBytes(vLog.Data)
		fmt.Println("vLog.Data(transferamount):", a)

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println("topics[0]:", vLog.Topics[0].Hex())
	}
	eventSignature := []byte("Transfer(address,address,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("method hash:", hash.Hex())
}
