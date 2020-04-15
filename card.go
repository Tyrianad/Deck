package main

import "fmt"

//Card rappresenta una carta da gioco
type Card struct {
	suit  string
	value string
}

func (c Card) print() {
	fmt.Print(c.value + " " + c.suit)
}

func (c Card) println() {
	fmt.Println(c.value + " " + c.suit)
}
