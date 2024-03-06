package main

//import "io/ioutil"

func main() {

	/*cards := newDeck()

	cards.saveToFile("my_cards")
	*/

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	deal(cards, 5)
	cards.print()
}
