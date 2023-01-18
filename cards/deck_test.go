package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expecting deck length of 16, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expecting first card to be Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expecting last card to be Four of Diamonds, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")
	d = newDeckFromFile("_deckTesting")
	if len(d) != 16 {
		t.Errorf("Expecting deck length of 16, got %v", len(d))
	}
	os.Remove("_deckTesting")
}
