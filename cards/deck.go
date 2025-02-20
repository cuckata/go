package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck'
//which is a slice of strings
type deck []string

//function with nested loop to create cards by Suit and Value combination
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", 
		"Eight", "Nine", "Ten", "Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards, value+ " of "+suit)
		}
	}
	
	return cards
}
//receiver function with for loop to print all cards in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
//deal function with arguments of deck and handSize, returns 
//user deck and the leftover cards in the main deck
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}
//receiver function to turn the deck into string
//so it can later be turned into byte slice
//in order to be saved to file
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
//saveToFile to create an actual physical file
//on the local machine
//using the type conversion from string to byte slice
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666 )
}
//import file from local machine into the app
//to create a new deck, with error handling
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		//log the error and entirely quit the program
		fmt.Println("error: ", err)
		os.Exit(1)
		//can also be made to log the error and return a call to NewDeck() 
		//so its still gives the user a deck to work with
	}
	//type conversion from byte slice to string and 
	//assigning the value to variable s
	s := strings.Split(string(bs), ",") 
	//create a deck from s and return it
	return deck(s)
}
//shuffle receiver function with deck as argument
//its not the best, but hey it works!!! no judging!
func (d deck) shuffle() {
	//since rand.Intn is pseudo-random number generator, we need
	//a way to make it properly randomize
	source := rand.NewSource(time.Now().UnixNano()) //create a source from current time converted into UnixNano which's value is of type int64
	r := rand.New(source) //use the source to create a new random value which can then generate other random values

	
	for i := range d {
		np:= r.Intn(len(d) - 1) //np is newPosition, r.Intn generates a random int in specified range
		
		d[i], d[np] = d[np], d[i] //swapping elements in deck slice
	}
}