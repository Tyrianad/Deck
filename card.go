package main

import "fmt"

//Card rappresenta una carta da gioco
type card struct {
	suit  string
	value string
}

func (c card) print() {
	fmt.Print(c.value + " " + c.suit)
}

func (c card) println() {
	fmt.Println(c.value + " " + c.suit)
}
