package main

func main() {
	cards := newDeck()
	cards.shuffle()
	dealt, remaining := deal(cards, 7)
	dealt.print()
	remaining.print()
}
