package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if d == nil {
		t.Error("Deck is not meant to be null")
	}

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	err := d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if err != nil {
		t.Errorf("Expected to run without errors, but got %v", err)
		os.Remove(filename)
		return
	}

	if loadedDeck == nil {
		t.Errorf("Expected to load deck from file, but recieved nil")
		os.Remove(filename)
		return
	}

	if len(d) != len(loadedDeck) {
		t.Errorf("Lenghts of initial deck and loaded deck are expected to match")
		os.Remove(filename)
		return
	}

	os.Remove(filename)
}
