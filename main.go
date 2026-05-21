package main

import "fmt"

// Fix: Use the 'var' keyword at the package level
var deckSize int = 20

// Alternatively, use 'const' if this value will never change:
// const deckSize = 20

func main() {
	var name string = "aakash"
	fmt.Println(name)
}
