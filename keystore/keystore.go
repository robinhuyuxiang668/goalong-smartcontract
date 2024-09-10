package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

/*
*
keystore是一个包含经过加密了的钱包私钥。go-ethereum中的keystore，每个文件只能包含一个钱包密钥对。
要生成keystore，首先您必须调用NewKeyStore，给它提供保存keystore的目录路径。然后，
您可调用NewAccount方法创建新的钱包，并给它传入一个用于加密的口令。您每次调用NewAccount，
它将在磁盘上生成新的keystore文件
*/
func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x14aC5219D35DA827Cc485309418d9eD05Bf58c92
}

func importKs() {
	file := "./tmp/UTC--2024-09-06T08-27-06.622914300Z--14ac5219d35da827cc485309418d9ed05bf58c92"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("ReadFile error:", err, "end")
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal("Import error:", err, " end")
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}
