package main

import (
	"math/rand"
)

//Deck rappresenta un mazzo di carte
type Deck struct {
	cards []Card
}

func (d Deck) printAll() {
	for _, card := range d.cards {
		card.println()
	}
}

func (d *Deck) shuffle() *Deck {
	var shuffledDeck Deck
	var j int
	r := rand.New(rand.NewSource(99))

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
func (d *Deck) AddCard(card Card) []Card {
	d.cards = append(d.cards, card)
	return d.cards
}

func (d *Deck) newDeck() *Deck {
	suits := []string{"♤", "♡", "♢", "♧"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			d.AddCard(Card{suit, value})
		}
	}

	return d
}
