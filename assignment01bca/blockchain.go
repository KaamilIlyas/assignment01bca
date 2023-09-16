package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// defining structure named "Block" to represent a single block in the blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// defining structure named "Blockchain" to represent the entire blockchain, which is a list of blocks
type Blockchain struct {
	Blocks []*Block
}

// creating & returning a new block with the given transaction data, nonce, and previous hash
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}

	//calculating and setting the current hash for the new block
	block.CurrentHash = CalculateHash(block)
	return block
}

// calculateHash calculates the hash of a block based on its transaction data, nonce, and previous hash
func CalculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// listBlocks prints information about each block in the blockchain
func (bc *Blockchain) ListBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i+1)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Current Hash: %s\n", block.CurrentHash)
		fmt.Println()
	}
}

// addBlock adds a new block to the blockchain with the given transaction data and nonce
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]                        //get the last block in the chain
	newBlock := NewBlock(transaction, nonce, previousBlock.CurrentHash) //create a new block
	bc.Blocks = append(bc.Blocks, newBlock)                             //append the new block to the blockchain
}

// changeBlock modifies the transaction and nonce of a specific block in the blockchain
func (bc *Blockchain) ChangeBlock(index int, newTransaction string, newNonce int) {
	if index >= 0 && index < len(bc.Blocks) {
		block := bc.Blocks[index]
		block.Transaction = newTransaction
		block.Nonce = newNonce
		block.CurrentHash = CalculateHash(block) //recalculate the hash after modification
	}
}

// verifyChain checks the integrity of the blockchain by verifying hashes and previous hash references
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false //if previous hash reference is incorrect, chain is invalid
		}

		if currentBlock.CurrentHash != CalculateHash(currentBlock) {
			return false //if current block's hash is incorrect, chain is invalid
		}
	}
	return true //if all checks pass, blockchain is considered valid
}
