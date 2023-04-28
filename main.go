package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := Block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash

	fmt.Println(genesisBlock)
}
