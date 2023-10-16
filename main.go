package main

import (
	"assignment01bca/assignment01bca"
	"fmt"
)

func main() {
	// creating a genesis block
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "", []*assignment01bca.Transaction{})

	// creating three nodes
	node1 := assignment01bca.NewNode("Node1")
	node2 := assignment01bca.NewNode("Node2")
	node3 := assignment01bca.NewNode("Node3")

	// adding the genesis block to each node's blockchain
	node1.Blockchain.AddBlock(genesisBlock)
	node2.Blockchain.AddBlock(genesisBlock)
	node3.Blockchain.AddBlock(genesisBlock)

	// simulating transactions for Node1
	node1.AddTransaction("Alice", "Bob", 1.0)
	node1.AddTransaction("Charlie", "Alice", 0.5)

	// mine blocks with difficulty level 2 for Node1
	node1.MineBlock(2)

	// printing blocks and transactions for Node1
	fmt.Println("Node 1 Blocks:")
	node1.ListBlocks()

	// verifying the nonce for the last mined block of Node1
	nonceToVerify := node1.Blockchain.Chain[len(node1.Blockchain.Chain)-1].Nonce
	node1.VerifyNonce(nonceToVerify)

	// simulating transactions for Node2
	node2.AddTransaction("Bob", "Charlie", 0.2)
	node2.AddTransaction("David", "Alice", 1.5)

	// mine blocks with difficulty level 2 for Node2
	node2.MineBlock(2)

	// printing blocks and transactions for Node2
	fmt.Println("Node 2 Blocks:")
	node2.ListBlocks()

	// verifying the nonce for the last mined block of Node2
	nonceToVerify = node2.Blockchain.Chain[len(node2.Blockchain.Chain)-1].Nonce
	node2.VerifyNonce(nonceToVerify)

	// simulate transactions for Node3
	node3.AddTransaction("Eve", "Frank", 0.8)

	// mine blocks with difficulty level 2 for Node3
	node3.MineBlock(2)

	// printing blocks and transactions for Node3
	fmt.Println("Node 3 Blocks:")
	node3.ListBlocks()

	// Verify the nonce for the last mined block of Node3
	nonceToVerify = node3.Blockchain.Chain[len(node3.Blockchain.Chain)-1].Nonce
	node3.VerifyNonce(nonceToVerify)
}
