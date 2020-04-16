package main

import (
	"math/rand"
	"time"
)

//deck rappresenta un mazzo di carte
type deck struct {
	cards []card
}

func (d deck) printAll() {
	for _, card := range d.cards {
		card.println()
	}
}

func (d *deck) shuffle() *deck {
	var shuffledDeck deck
	var j int
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

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

func (d *deck) drawN(nCards int) []card {
	//@TODO gestire il drawN su deck con meno carte
	c := d.cards[:nCards]
	d.cards = d.cards[nCards:]

	return c
}
