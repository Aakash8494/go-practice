package main

import (
	"fmt"
	"runtime"
)

// Fix: Use the 'var' keyword at the package level
// deckSize := 20
// var deckSize int = 20

func main() {
	// var name string = "aakash"
	// fmt.Println(name)
	// name = "aakash1"
	// fmt.Println(name)

	// name1 := "harshada"
	// fmt.Println(name1)
	// name1 = "harshada1"
	// fmt.Println(name1)

	// card1 := newCard()
	// fmt.Println(card1)

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	// This will ask your machine exactly how many cores it has!
	totalBrainPower := runtime.NumCPU()
	fmt.Printf("Bro, your laptop has %d CPU cores!\n", totalBrainPower)
}

func newCard() string {
	return "Ace of Spades"
}
