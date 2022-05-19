package main

import (
	"jenkka/cards/deck"
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := deck.NewDeck()

	if len(newDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(newDeck))
	}

	FIRST_CARD_NAME := "Ace of Clubs"
	if newDeck[0] != FIRST_CARD_NAME {
		t.Errorf("Expected first card to be %s, but got %s", FIRST_CARD_NAME, newDeck[0])
	}

	LAST_CARD_NAME := "King of Diamonds"
	if newDeck[len(newDeck)-1] != LAST_CARD_NAME {
		t.Errorf("Expected last card to be %s, but got %s", LAST_CARD_NAME, newDeck[len(newDeck)-1])
	}
}

func TestSaveAndLoadFile(t *testing.T) {
	TEST_FILE_NAME := "_decktesting"

	// Cleaning environment
	os.Remove(TEST_FILE_NAME)

	newDeck := deck.NewDeck()
	newDeck.SaveToFile(TEST_FILE_NAME)

	loadedDeck := deck.NewDeckFromFile(TEST_FILE_NAME)

	if !reflect.DeepEqual(newDeck, loadedDeck) {
		t.Errorf("Saved and loaded deck are not equal. Fix functions.")
	}

	os.Remove(TEST_FILE_NAME)
}
