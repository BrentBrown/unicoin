package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	sender    string
	recipient string
	value     float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.sender,
		Recipient: t.recipient,
		Value:     t.value,
	})
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address      %s\n", t.sender)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.recipient)
	fmt.Printf(" value                          %.1f\n", t.value)
}
