package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//deck rappresenta un mazzo di carte
type deck struct {
	cards []card
}

func (d deck) print() {
	fmt.Println(d.toString())
}

func (d *deck) shuffle() *deck {
	var shuffledDeck deck
	var j int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	totalCards := len(d.cards)

	for i := 0; i < totalCards; i++ {
		j = r.Intn(len(d.cards))
		shuffledDeck.AddCard(d.cards[j])
		d.cards = append(d.cards[:j], d.cards[j+1:]...)
	}

	d.cards = shuffledDeck.cards

	return d
}

//AddCard permette di aggiungere una carta al mazzo
func (d *deck) AddCard(card card) []card {
	d.cards = append(d.cards, card)
	return d.cards
}

func (d *deck) newDeck() *deck {
	suits := []string{"♤", "♡", "♢", "♧"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			d.AddCard(card{suit, value})
		}
	}

	return d
}

func (d *deck) draw() card {
	//@TODO gestire il draw su deck vuoto
	c := d.cards[0]
	d.cards = d.cards[1:]

	return c
}

func (d *deck) deal(n int) []card {
	//@TODO gestire il deal su deck con meno carte
	c := d.cards[:n]
	d.cards = d.cards[n:]

	return c
}

func (d *deck) toString() string {
	var ret string
	for _, card := range d.cards {
		ret = ret + "|" + card.toString()
	}

	ret = strings.TrimLeft(ret, "|")

	return ret
}

//deckFromString returns a deck based on incoming string configuration. Incoming string should be in the "toString" format.
func deckFromString(st string) deck {
	var d deck
	cards := strings.Split(st, "|")
	for _, card := range cards {
		d.AddCard(cardFromString(card))
	}
	return d
}

func (d *deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func (d *deck) loadFromFile(filename string) *deck {
	//b:= byteslice coming from readfile, contains the converted deck
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	tempDeck := deckFromString(string(b))

	d.cards = tempDeck.cards

	return d
}
