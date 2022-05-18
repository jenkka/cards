package deck

import "fmt"

type deck []string

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func NewDeck() deck {
	suits := [4]string{"Clubs", "Hearts", "Spades", "Diamonds"}
	numbers := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	deck := deck{}

	for _, suit := range suits {
		for _, number := range numbers {
			newCard := number + " of " + suit
			deck = append(deck, newCard)
		}
	}

	return deck
}
