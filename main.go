package main

import (
	"jenkka/cards/deck"
)

func main() {
	myDeck := deck.NewDeck()
	// myDeck.Shuffle()

	// fmt.Println("FIRST HAND")
	// myHand := myDeck.Deal(5)
	// myHand.Print()

	// fmt.Println("SECOND HAND")
	// secondHand := myDeck.Deal(10)
	// secondHand.Print()
	// fmt.Println("---------------")
	myDeck.Print()
}
