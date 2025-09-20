package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timeStamp    int64
	transaction  []string
}

func NewBlock(nonce int, previousHash string) *Block {
	block := new(Block)
	block.timeStamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash

	return block
}

func (block *Block) PrintBlock() {
	fmt.Printf("TIMESTAMP : %d", block.timeStamp)
	fmt.Printf("\nPREVIOUS HASH : %v", block.previousHash)
	fmt.Printf("\nNONCE : %d", block.nonce)
	fmt.Printf("\nTRANSACTION : %v", block.transaction)
}

func init() {
	log.SetPrefix("Blockchain Node : ")
}

func main() {
	block := NewBlock(0, "x000")
	block.PrintBlock()

}
