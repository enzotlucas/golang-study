package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("------------------------")
	fmt.Println("new:")
	fmt.Println("------------------------")
	cards.shufle()
	cards.print()
}
