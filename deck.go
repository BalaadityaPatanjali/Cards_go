package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for i := 0; i < len(cardSuits); i++ {
		for j := 0; j < len(cardValues); j++ {
			cards = append(cards, cardValues[i]+" of "+cardSuits[j])
		}
	}
	return cards
}

func (d deck) print() {
	for i := 0; i < len(d); i++ {
		fmt.Println(i, d[i])
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckfromfile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < len(d); i++ {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
