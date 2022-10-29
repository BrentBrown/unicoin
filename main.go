package main

func main() {
	bc := NewBlockchain()
	//bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	ph := bc.LastBlock().Hash()
	bc.CreateBlock(5, ph)
	//bc.Print()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	ph = bc.LastBlock().Hash()
	bc.CreateBlock(2, ph)
	bc.Print()
}
