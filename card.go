package main

import (
	"fmt"
	"strings"
)

//Card rappresenta una carta da gioco
type card struct {
	suit  string
	value string
}

func (c card) print() {
	fmt.Println(c.toString())
}

func (c card) toString() string {
	return c.value + "_" + c.suit
}

func cardFromString(st string) card {
	s := strings.Split(st, "_")

	c := card{suit: s[1], value: s[0]}

	return c
}
