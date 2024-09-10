package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 用FromECDSA方法将其转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//得到私钥 : 0X开头
	fmt.Println(hexutil.Encode(privateKeyBytes))
	//得到私钥 : 去掉0X开头
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//得到公钥
	fmt.Println(hexutil.Encode(publicKeyBytes))
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	//得到地址 0x96216849c49358B10257cb55b28eA603c874b05E
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	//另一种由公钥推出地址方式:(公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）
	//并用“0x”作为前缀。 以下是使用 golang.org/x/crypto/sha3 的 Keccak256函数手动完成的方法)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
