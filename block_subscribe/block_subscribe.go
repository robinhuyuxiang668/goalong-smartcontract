package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/GBDdpIse13KF_izTI2-J5KACwNNtxS0Z")
	if err != nil {
		log.Fatal(err)
	}

	//创建一个新的通道，用于接收最新的区块头
	headers := make(chan *types.Header)
	//传入的第二个参数使 headers变为只接收的通道,用于在下面从通道读取最新header信息
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		//从通道读取
		case header := <-headers:
			fmt.Println("new header hash:", header.Hash().Hex()) //header Hash和区块hash一样

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
