package main

func main() {
	deck := Deck{}
	deck.newDeck()
	//deck.printAll()
	deck.shuffle()
	deck.printAll()
}
