package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce        int      `json:"nonce"`
	PreviousHash [32]byte `json:"previousHash"`
	Timestamp    int64    `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := b.marshalJSON()
	return sha256.Sum256(m)
}

func (b *Block) marshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Block) Print() {
	fmt.Printf("Timestamp       %d\n", b.Timestamp)
	fmt.Printf("Nonce           %d\n", b.Nonce)
	fmt.Printf("previous_hash   %x\n", b.PreviousHash)
	fmt.Printf("Transactions    %s\n", b.Transactions)
}
