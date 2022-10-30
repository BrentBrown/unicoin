package main

import (
	"fmt"
	"log"
	"unicoin/wallet"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
