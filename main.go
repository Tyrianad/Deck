package main

func main() {
	deck := deck{}
	deck.newDeck()
	deck.shuffle()
	deck.loadFromFile("test")
	deck.print()
}
