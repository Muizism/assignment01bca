package main

import (
    "crypto/sha256"
    "fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    Hash         string
}

// Blockchain represents the entire blockchain.
type Blockchain struct {
    Blocks []*Block
}

// NewBlock creates a new block with the given data.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
    block := &Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
    }
    block.Hash = CalculateHash(transaction, nonce, previousHash)
    return block
}

// DisplayBlocks prints all blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
    for i, block := range bc.Blocks {
        fmt.Printf("Block %d:\n", i)
        fmt.Printf("  Transaction: %s\n", block.Transaction)
        fmt.Printf("  Nonce: %d\n", block.Nonce)
        fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("  Current Hash: %s\n", block.Hash)
        fmt.Println()
    }
}

// ChangeBlock changes the transaction data of a given block.
func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
    if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
        bc.Blocks[blockIndex].Transaction = newTransaction
        bc.Blocks[blockIndex].Hash = CalculateHash(newTransaction, bc.Blocks[blockIndex].Nonce, bc.Blocks[blockIndex].PreviousHash)
    }
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
    for i := 1; i < len(bc.Blocks); i++ {
        currentBlock := bc.Blocks[i]
        previousBlock := bc.Blocks[i-1]
        if currentBlock.Hash != CalculateHash(currentBlock.Transaction, currentBlock.Nonce, previousBlock.Hash) {
            return false
        }
    }
    return true
}

// CalculateHash calculates the hash of a given string.
func CalculateHash(stringToHash string, nonce int, previousHash string) string {
    data := fmt.Sprintf("%s%d%s", stringToHash, nonce, previousHash)
    hash := sha256.Sum256([]byte(data))
    return fmt.Sprintf("%x", hash)
}

func main() {
    // Create a new blockchain with the genesis block.
    genesisBlock := NewBlock("Genesis Block", 0, "")
    blockchain := &Blockchain{Blocks: []*Block{genesisBlock}}

    // Add some more blocks to the blockchain.
    blockchain.Blocks = append(blockchain.Blocks, NewBlock("Alice to Bob", 123, blockchain.Blocks[len(blockchain.Blocks)-1].Hash))
    blockchain.Blocks = append(blockchain.Blocks, NewBlock("Bob to Carol", 456, blockchain.Blocks[len(blockchain.Blocks)-1].Hash))

    // Display all blocks in the blockchain.
    blockchain.DisplayBlocks()

    // Change the transaction data of the second block.
    blockchain.ChangeBlock(1, "New Transaction Data")

    // Verify the blockchain.
    if blockchain.VerifyChain() {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain is not valid.")
    }
}
