package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://arb-sepolia.g.alchemy.com/v2/GBDdpIse13KF_izTI2-J5KACwNNtxS0Z")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x7C44AfdCcc6857338bd30CE9CD0C6c3788309777")
	//设置监听事件的合约地址
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	//接收事件发出的LOG
	logs := make(chan types.Log)
	//订阅 监听后面所有指定地址EVENT
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
