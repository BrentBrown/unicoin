package main

func main() {
	bc := NewBlockchain()
	bc.Print()

	ph := bc.LastBlock().Hash()
	bc.CreateBlock(5, ph)
	bc.Print()

	ph = bc.LastBlock().Hash()
	bc.CreateBlock(2, ph)
	bc.Print()
}
