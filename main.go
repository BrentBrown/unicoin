package main

func main() {
	bc := NewBlockchain()
	//bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	ph := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, ph)
	//bc.Print()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	ph = bc.LastBlock().Hash()
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce, ph)
	bc.Print()
}
