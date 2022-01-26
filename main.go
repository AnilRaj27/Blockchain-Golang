package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Data          []byte
	BlockHash     []byte
	PrevBlockHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(info)
	b.BlockHash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		BlockHash:     []byte{},
		Data:          []byte(data),
		PrevBlockHash: prevHash,
	}
	block.DeriveHash()

	return block
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.BlockHash)
	chain.blocks = append(chain.blocks, new)
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("1st Block")
	chain.AddBlock("2nd Block")
	chain.AddBlock("3rd Block")

	for _, block := range chain.blocks {
		fmt.Printf("Hash: %x\n", block.BlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("\n")
	}
}
