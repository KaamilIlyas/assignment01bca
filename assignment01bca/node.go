package assignment01bca

import (
	"fmt"
	"sync"
)

// node structure
type Node struct {
	Blockchain *Blockchain
	NodeName   string
	mutex      sync.Mutex //mutex to ensure safe concurrent access to the node's data
}

// creating a new node
func NewNode(nodeName string) *Node {
	return &Node{
		Blockchain: &Blockchain{Chain: []*Block{}},
		NodeName:   nodeName,
	}
}

// it verifies if the given nonce is the correct one for the last mined block
// it calculates derived nonce for last mined block using DeriveNonce function and appropriate difficulty level
func (n *Node) VerifyNonce(nonceToVerify int) {

	lastMinedBlock := n.Blockchain.Chain[len(n.Blockchain.Chain)-1]
	derivedNonce := DeriveNonce(lastMinedBlock, 2)

	// Compare the derived nonce with the nonce to verify
	if nonceToVerify == derivedNonce {
		fmt.Printf("Nonce %d is correct for the last mined block.\n\n\n", nonceToVerify)
	} else {
		fmt.Printf("Nonce %d is incorrect for the last mined block. Expected nonce: %d\n\n\n", nonceToVerify, derivedNonce)
	}

}

// adding a Transaction to the transaction pool
// transaction represents a transfer of cryptocurrency from one blockchain address to another with a specified value
func (n *Node) AddTransaction(sender string, recipient string, value float32) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	t := NewTransaction(sender, recipient, value)
	n.Blockchain.TransactionPool = append(n.Blockchain.TransactionPool, t)
}

// it is a key operation where the node competes to find a valid nonce for the new block
func (n *Node) MineBlock(difficulty int) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	minedBlock := n.Blockchain.MineBlock(difficulty)
	n.Blockchain.AddBlock(minedBlock)
}

// it lists the blocks stored in the node's blockchain
func (n *Node) ListBlocks() {
	fmt.Printf("Node %s Blocks:\n", n.NodeName)
	n.Blockchain.ListBlocks()
}
