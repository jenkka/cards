package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d))
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func (d *deck) Deal(handSize int) deck {
	hand := (*d)[:handSize]
	*d = (*d)[handSize:]

	return hand
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	deck := deck(strings.Split(string(byteSlice), ","))

	return deck
}
