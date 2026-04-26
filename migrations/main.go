package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 1. 生成测试私钥（仅用于测试）
	//privateKey, err := crypto.GenerateKey()
	privateKey, err := crypto.HexToECDSA("e4c39de79c1c67c93f5914a479c8471f9e0a7caf9e58c42dbc9f3c63f3202679")
	if err != nil {
		panic(err)
	}

	// 2. 要签名的消息
	message := `Welcome to Coinroll!\n\nClick to sign in and accept the Terms of Service.\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nWallet address:\n0x1234567890123456789012345678901234567890\n\nNonce:\n7c5abda8af86309c2f3010d4975bb3f0`

	// 3. 添加以太坊签名前缀并签名
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	prefixedMessage := prefix + message

	signature, err := crypto.Sign(crypto.Keccak256Hash([]byte(prefixedMessage)).Bytes(), privateKey)
	if err != nil {
		panic(err)
	}

	// 4. 输出结果
	fmt.Println("PrivateKey:", crypto.FromECDSA(privateKey))
	fmt.Println("Address:", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	fmt.Println("Signature:", fmt.Sprintf("0x%x", signature))
}
