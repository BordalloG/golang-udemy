package main

func main() {
	cards := newDeck()
	cards.saveToFile("teste.txt")
	newDeckFromFile("teste.txt").print()
}
