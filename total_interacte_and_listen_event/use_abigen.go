package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"sync"
)

const (
	RPC             = "https://eth-sepolia.g.alchemy.com/v2/GBDdpIse13KF_izTI2-J5KACwNNtxS0Z"
	WC              = "wss://eth-sepolia.g.alchemy.com/v2/GBDdpIse13KF_izTI2-J5KACwNNtxS0Z"
	privateKey      = ""
	toAddress       = ""
	contractAddress = "0x7697B3f8AF0fC50E58452C1f39AA67D97a9909B4"
)

var wg sync.WaitGroup
var cAddr common.Address

func main() {
	client, err := ethclient.Dial(RPC)
	if err != nil {
		fmt.Println("ethclient.Dial error : ", err)
		os.Exit(1)
	}

	cAddr = common.HexToAddress(contractAddress)
	yolo, err := NewYolo(cAddr, client)
	if err != nil {
		fmt.Println("NewYolo error : ", err)
		return
	}

	wg.Add(1)
	go listenEvent()

	lastRound, err := yolo.LastRound(nil)
	if err != nil {
		fmt.Println("LastRound error:", err)
		return
	}
	fmt.Println("LastRound:", lastRound)

	// 获取当前区块链的ChainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("获取ChainID失败:", err)
		return
	}
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("privateKeyECDSA error ,", err)
		return
	}

	//构建参数对象
	opts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		fmt.Println("bind.NewKeyedTransactorWithChainID error ,", err)
		return
	}

	bets := []YoloVRFBet{
		YoloVRFBet{3, 30},
		YoloVRFBet{7, 18},
		YoloVRFBet{8, 15},
		YoloVRFBet{1, 13},
		YoloVRFBet{2, 13},
		YoloVRFBet{4, 9},
		YoloVRFBet{5, 5},
		YoloVRFBet{10, 5},
		YoloVRFBet{9, 4},
		YoloVRFBet{6, 4},
	}
	lastRound++
	tx, err := yolo.Betting(opts, lastRound, bets)
	if err != nil {
		fmt.Println("Betting error:", err)
		return
	}
	fmt.Println("tx:", tx)

	//delay := 20 * time.Second
	//fmt.Println("延时开始前")
	//time.Sleep(delay) // 执行延时
	//fmt.Println("延时结束后")
	//
	//betsNew, err := yolo.BetInfo(nil, lastRound)
	//if err != nil {
	//	fmt.Println("BetInfo error:", err)
	//	return
	//}
	//fmt.Println("bets:", betsNew)
	//
	//winner, err := yolo.Winner(nil, lastRound)
	//if err != nil {
	//	fmt.Println("BetInfo error:", err)
	//	return
	//}
	//fmt.Println("winner:", winner)
	wg.Wait()
}

func listenEvent() {
	client, err := ethclient.Dial(WC)
	if err != nil {
		fmt.Println("listenEvent ethclient.Dial error : ", err)
		os.Exit(1)
	}

	//设置监听事件的合约地址
	query := ethereum.FilterQuery{
		Addresses: []common.Address{cAddr},
	}

	//接收事件发出的LOG
	logs := make(chan types.Log)
	//订阅 监听后面所有指定地址EVENT
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	eventSignature := []byte("Win(uint64,uint64)")
	hash := crypto.Keccak256Hash(eventSignature)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			fmt.Println("BlockHash ", vLog.BlockHash.Hex())
			fmt.Println("BlockNumber ", vLog.BlockNumber)
			fmt.Println("TxHash ", vLog.TxHash.Hex())
			a := new(big.Int)
			a.SetBytes(vLog.Data)
			fmt.Println("vLog.Data:", a)

			//过滤win事件
			if len(vLog.Topics) == 3 && hash.Hex() == vLog.Topics[0].Hex() {
				round := vLog.Topics[1].Big()
				winId := vLog.Topics[2].Big()
				fmt.Printf("round:%v ,winId:%v\n", round, winId)
			}
			//wg.Done()
			//break
		}
	}

}
