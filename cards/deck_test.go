package main

import (
	"os"
	"testing"
)

//test that deck is with 52 number of cards
func TestNewDeckCreation(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

//test that deck starts with Ace of Spades card and ends with King of Clubs card,
//due to the logic of the for loops in newDeck()
func TestNewDeckContents(t *testing.T) {
	d :=newDeck()

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v instead", d[0])
	}
	if d[len(d) - 1] != "King of Clubs"{
		t.Errorf("Expected last card to be 'King of Clubs', but got %v instead", d[len(d) - 1])
	}
}
//test that shuffle works,
//not really sure how to test it other than
//to check first two cards and last two cards
//and verify they are not from same suit and value+1
//and yes im aware that there is still a minor chance
//that even after shuffle the test can fail which makes it flaky
//but i dont wanna spend more time on this right now
func TestShuffle(t *testing.T){
	d :=newDeck()
	d.shuffle()

	if d[0] == "Ace of Spades" && d[1] == "Two of Spades" {
		t.Errorf("First card is 'Ace of Spades' and second is 'Two of spades', which might be just a coincidence, but check shuffle() logic")
	}

	if d[50] == "Queen of Clubs" && d[51] == "King of Clubs" {
		t.Errorf("51st card is 'Queen of Clubs' and second is 'King of Clubs', which might be just a coincidence, but check shuffle() logic")	
	}
}
//test saveToFile and newDeckFromFile
//yes im aware the test name is ridiculous
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	//first we try to remove an old file with the same name
	//just in case in alst run we had a crash in between
	//creation of file and loading of file
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}