package main

import "fmt"

func main() {
	myAddress := "myAddress"
	bc := NewBlockchain(myAddress)
	//bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	//bc.Print()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	bc.Mining()
	bc.Print()

	fmt.Printf("my %.1f\n", bc.CalculateTotalAmount(myAddress))
	fmt.Printf("C %.1f\n", bc.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", bc.CalculateTotalAmount("D"))
}
