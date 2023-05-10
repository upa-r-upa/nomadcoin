package main

import (
	"fmt"

	"github.com/upa-r-upa/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
