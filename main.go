package main

import "fmt"

func main() {
	deck := deck{}
	deck.newDeck()
	deck.shuffle()
	deck.loadFromFile("test")
	fmt.Println(deck)
}
