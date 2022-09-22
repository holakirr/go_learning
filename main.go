package main

func main() {
	// var card string = "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	// fmt.Println(card)

	cards := deck{"Ace of Spades", newCard(), "Ace of Diamond"}
	cards = append(cards, "Six of Diamonds")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
