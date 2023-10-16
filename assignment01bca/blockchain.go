package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// representing a block in the blockchain
type Block struct {
	Transaction     string
	Nonce           int
	PreviousHash    string
	CurrentHash     string
	TransactionPool []*Transaction
	Timestamp       int64
}

// creating a new block
func NewBlock(transaction string, nonce int, previousHash string, transactionPool []*Transaction) *Block {
	block := &Block{
		Transaction:     transaction,
		Nonce:           nonce,
		PreviousHash:    previousHash,
		TransactionPool: transactionPool,
		Timestamp:       time.Now().Unix(),
	}

	//calculate and set the current hash for the new block
	block.CurrentHash = CalculateHash(block.Transaction + strconv.Itoa(nonce) + block.PreviousHash)
	return block
}

// calculates the SHA-256 hash of a given string
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

// defining Transaction structure
type Transaction struct {
	TransactionID              string
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

// creating a new Transaction and calculating a unique transaction ID based on the inputs
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	transactionID := CalculateHash(sender + recipient + fmt.Sprint(value))
	return &Transaction{
		TransactionID:              transactionID,
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
}

// defining Blockchain structure
type Blockchain struct {
	Chain           []*Block
	TransactionPool []*Transaction
}

// adding a Block to the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	bc.Chain = append(bc.Chain, block)
}

// listBlocks lists blocks with transactions in JSON format
func (bc *Blockchain) ListBlocks() {
	for i, block := range bc.Chain {
		fmt.Printf("\n=============== Block %d: ===============\n", i+1)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Timestamp: %s\n", time.Unix(block.Timestamp, 0))
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Current Hash: %s\n", block.CurrentHash)
		fmt.Println("--------------------------------")
		fmt.Println("  Transactions:")
		bc.ListTransactionsOfBlockInJSON(i) // Display transactions in JSON format
		fmt.Println()
	}
}

// listTransactionsOfBlock lists transactions of a specific block
func (bc *Blockchain) ListTransactionsOfBlock(blockIndex int) {
	if blockIndex >= 0 && blockIndex < len(bc.Chain) {
		block := bc.Chain[blockIndex]
		fmt.Printf("Transactions in Block %d:\n", blockIndex+1)
		for i, transaction := range block.TransactionPool {
			fmt.Printf("Transaction %d:\n", i+1)
			fmt.Printf("  Transaction ID: %s\n", transaction.TransactionID)
			fmt.Printf("  Sender: %s\n", transaction.SenderBlockchainAddress)
			fmt.Printf("  Recipient: %s\n", transaction.RecipientBlockchainAddress)
			fmt.Printf("  Value: %f\n", transaction.Value)
			fmt.Println()
		}
	}
}

// listTransactionsOfBlockInJSON lists transactions of a specific block in JSON format
func (bc *Blockchain) ListTransactionsOfBlockInJSON(blockIndex int) {
	if blockIndex >= 0 && blockIndex < len(bc.Chain) {
		block := bc.Chain[blockIndex]
		transactions := block.TransactionPool
		jsonTransactions, err := json.Marshal(transactions)
		if err != nil {
			fmt.Println("Error marshaling transactions into JSON:", err)
			return
		}
		fmt.Printf("Transactions in Block %d (in JSON format):\n%s\n", blockIndex+1, jsonTransactions)
	}
}

// function to find nonce that satisfies POW requirement (a prefix of zeros) for a given block and difficulty level
func DeriveNonce(block *Block, difficulty int) int {
	prefix := strings.Repeat("0", difficulty)
	nonce := 0
	for {
		hash := CalculateHash(block.Transaction + strconv.Itoa(nonce) + block.PreviousHash)
		if strings.HasPrefix(hash, prefix) {
			return nonce
		}
		nonce++
	}
}

// method to mine a new block with a specified difficulty level
// it selects the previous block's hash (if available) as the previous hash and includes transactions from the transaction pool
func (bc *Blockchain) MineBlock(difficulty int) *Block {
	previousHash := ""
	if len(bc.Chain) > 0 {
		previousHash = bc.Chain[len(bc.Chain)-1].CurrentHash
	}

	transaction := ""
	for _, t := range bc.TransactionPool {
		transaction += t.TransactionID
	}

	block := NewBlock(transaction, 0, previousHash, bc.TransactionPool)

	//block is then mined using POW with difficulty and mined block is added to blockchain
	nonce := DeriveNonce(block, difficulty)
	block.Nonce = nonce
	block.CurrentHash = CalculateHash(block.Transaction + strconv.Itoa(nonce) + block.PreviousHash)
	bc.TransactionPool = []*Transaction{} //clearing the transaction pool after mining
	return block
}
