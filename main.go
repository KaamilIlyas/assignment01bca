package main

import (
	"assignment01bca/assignment01bca"
	"fmt"
)

func main() {

	//creating a new blockchain with an initial block (the Genesis Block)
	initialBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	blockchain := &assignment01bca.Blockchain{Blocks: []*assignment01bca.Block{initialBlock}}

	//adding blocks to the blockchain
	blockchain.AddBlock("Transaction 1", 123)
	blockchain.AddBlock("Transaction 2", 456)

	//listing all blocks in the blockchain
	blockchain.ListBlocks()

	//tamper with a block by changing the transaction data
	blockchain.ChangeBlock(1, "Modified Transaction", 789)

	//list all blocks again
	blockchain.ListBlocks()

	//verifying the integrity of the blockchain
	isValid := blockchain.VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is invalid.")
	}
}
