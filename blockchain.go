package main

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timeStamp    int64
	transaction  []string
}

type Blockchain struct {
	transactionPool  []string
	transactionChain []*Block
}

func NewBlock(nonce int, previousHash string) *Block {
	block := new(Block)
	block.timeStamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash

	return block
}

func (blockchain *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	block_ := NewBlock(nonce, previousHash)
	blockchain.transactionChain = append(blockchain.transactionChain, block_)
	return block_
}

func NewBlockchain() *Blockchain {
	blockchain := new(Blockchain)
	genesisBlock := blockchain.CreateBlock(0, "x000")
	blockchain.transactionChain = append(blockchain.transactionChain, genesisBlock)
	return blockchain
}

func (block *Block) PrintBlock() {
	fmt.Printf("TIMESTAMP : %d", block.timeStamp)
	fmt.Printf("\nPREVIOUS HASH : %v", block.previousHash)
	fmt.Printf("\nNONCE : %d", block.nonce)
	fmt.Printf("\nTRANSACTION : %v", block.transaction)
}

func (blockchain *Blockchain) PrintBlockchain() {
	fmt.Printf("CURRENT TRANSACTION POOL : %v", blockchain.transactionPool)
	fmt.Printf("CURRENT CHAIN : %v", blockchain.transactionChain)
}

// func init() {
// 	log.SetPrefix("Blockchain Node : ")
// }

func main() {
	blockchain := NewBlockchain()
	blockchain.PrintBlockchain()
}
