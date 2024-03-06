package main

//import "io/ioutil"

func main() {

	/*cards := newDeck()

	cards.saveToFile("my_cards")
	*/

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	ha, rem := deal(cards, 5)

	ha.print()
	rem.print()

}
