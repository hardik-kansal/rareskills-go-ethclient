package main

import (
	"fmt"
)

const (
	sepoliaRpcUrl = "https://eth-sepolia.g.alchemy.com/v2/GNp-Xq9PjB97HhJo1H02Y-qpBHCHXOeO" // sepolia rpc url
	/* mainnetRpcUrl = "https://rpc.builder0x69.io/"   */ // mainnet rpc url
	from= "0x63f3405E55D2d209F0B89F6E6D7c077409b30a7C"
	to= "0x032ADe3f3e274Ed2171A5C099929616817Dc59e2"
	data= "Hello Ethereum!"
	privKey= ""
	gasLimit= uint64(21500) // adjust this if necessary
	wei= uint64(0)     // 0 Wei
)

func main() {
	fmt.Println("using ethclient...")

	getSuggestedGasPrice(sepoliaRpcUrl)

	eGas := estimateGas(sepoliaRpcUrl, from, to, data, wei)
	fmt.Println("\nestimate gas for the transaction is:", eGas)

	rawTxRLPHex := createRawTransaction(sepoliaRpcUrl, to, data, privKey, gasLimit, wei)
	fmt.Println("\nRaw TX:\n", rawTxRLPHex)
	sendRawTransaction(rawTxRLPHex, sepoliaRpcUrl)

	sig, sDetails := signMessage(data, privKey)
	fmt.Println("\nsigned message:", sDetails)

	if isSigner := verifySig(sig, from, data); isSigner {
		fmt.Printf("\n%s signed %s\n", from, data)
	} else {
		fmt.Printf("\n%s did not sign %s\n", from, data)
	}

	cNonce, nNonce := getNonce(to, sepoliaRpcUrl)
	fmt.Printf("\n%s current nonce: %v\n", to, cNonce)
	fmt.Printf("%s next nonce: %v\n", to, nNonce)

/* 	res := traceTx("0xd12e31c3274ff32d5a73cc59e8deacbb0f7ac4c095385add3caa2c52d01164c1", sepoliaRpcUrl)
	fmt.Println("\ntrace result:\n", res)

	blob, err := sendBlobTX(sepoliaRpcUrl, to, data, privKey)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nBlob transaction hash:", blob) */
}
