package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
	var value int
	switch card {
	case "ace":
		value = 11
	case "king", "queen", "jack", "ten":
		value = 10
	case "nine":
		value = 9
	case "eight":
		value = 8
	case "seven":
		value = 7
	case "six":
		value = 6
	case "five":
		value = 5
	case "four":
		value = 4
	case "three":
		value = 3
	case "two":
		return 2
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
	var value string
	switch {
	case ParseCard(card1) == 11 && ParseCard(card2) == 11:
		value = "P"
	// case ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) != 11 && ParseCard(dealerCard) != 10:
	case ParseCard(card1)+ParseCard(card2) == 21:
		// value = "W"
		if ParseCard(dealerCard) == 11 || ParseCard(dealerCard) == 10 {
			value = "S"
		}
		if ParseCard(dealerCard) != 11 && ParseCard(dealerCard) != 10 {
			value = "W"
		}
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		value = "S"
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		value = "S"
		if ParseCard(dealerCard) >= 7 {
			value = "H"
		}
	case ParseCard(card1)+ParseCard(card2) <= 11:
		value = "H"
	}
	return value
}
