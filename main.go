package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) <= 0 {
		return ""
	}

	return b.blocks[len(b.blocks)-1].hash
}

func (b *blockchain) addBlock(data string) {
	nextBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(nextBlock.data + nextBlock.prevHash))
	nextBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, nextBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("PrevHash: %s\n", block.prevHash)
		fmt.Println("---------------")
	}
}

func main() {
	chain := blockchain{}

	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}
