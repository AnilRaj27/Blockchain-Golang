package main

import (
	"fmt"
	"strconv"

	"github.com/AnilRaj27/Blockchain-Golang/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("1st Block")
	chain.AddBlock("2nd Block")
	chain.AddBlock("3rd Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash: %x\n", block.BlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("\n")

		pow := blockchain.NewProof(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	// target := big.NewInt(1)
	// fmt.Println(target)
	// fmt.Println(target.Lsh(target, 257))

}
