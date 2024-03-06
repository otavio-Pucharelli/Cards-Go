package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Create a deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a new function for the type 'deck'
func (d deck) print() { //* (d deck) is a receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// []string(d) converts the deck to a slice of strings
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is a file permission
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)

		// os.Exit(1)
		return newDeck()
	}

	s := strings.Split(string(bs), ",") // convert byte slice to string and then to slice of strings
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		// generate a random number between 0 and length of slice - 1
		newPostion := rand.Intn(len(d) - 1)

		// swap the elements
		d[i], d[newPostion] = d[newPostion], d[i]
	}
}
