package main

func main() {
	cards := newDeckFromFile("deck1.txt")
	cards.print()
	cards.shuffle()
	cards.print()
}
