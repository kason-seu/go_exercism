package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.

var card2Int = map[string]int{
	"ace":   11,
	"eight": 8,
	"two":   2,
	"nine":  9,
	"three": 3,
	"ten":   10,
	"four":  4,
	"jack":  10,
	"five":  5,
	"queen": 10,
	"six":   6,
	"king":  10,
	"seven": 7,
	"other": 0,
}

func ParseCard(card string) int {
	return card2Int[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := card2Int[card1] + card2Int[card2]
	dealerCardInt := card2Int[dealerCard]
	switch {
	case sum == 22:
		return "P"
	case sum == 21 && dealerCardInt < 10:
		return "W"
	case sum == 21 && dealerCardInt > 10:
		return "S"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardInt < 7:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardInt >= 7:
		return "H"
	case sum <= 11:
		return "H"
	default:
		return "S"
	}

}
