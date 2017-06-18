package main

import (
	"fmt"
	"strconv"
)

type PaiValue int

const (
	Dots1 PaiValue = iota
	Dots2
	Dots3
	Dots4
	Dots5
	Dots6
	Dots7
	Dots8
	Dots9
	Bamboo1
	Bamboo2
	Bamboo3
	Bamboo4
	Bamboo5
	Bamboo6
	Bamboo7
	Bamboo8
	Bamboo9
	Characters1
	Characters2
	Characters3
	Characters4
	Characters5
	Characters6
	Characters7
	Characters8
	Characters9
	East
	South
	West
	North
	Red
	Green
	White
)

func (p PaiValue) String() string {
	switch p {
	case Dots1:
		return "Dots1"
	case Dots2:
		return "Dots2"
	case Dots3:
		return "Dots3"
	case Dots4:
		return "Dots4"
	case Dots5:
		return "Dots5"
	case Dots6:
		return "Dots6"
	case Dots7:
		return "Dots7"
	case Dots8:
		return "Dots8"
	case Dots9:
		return "Dots9"
	case Bamboo1:
		return "Bamboo1"
	case Bamboo2:
		return "Bamboo2"
	case Bamboo3:
		return "Bamboo3"
	case Bamboo4:
		return "Bamboo4"
	case Bamboo5:
		return "Bamboo5"
	case Bamboo6:
		return "Bamboo6"
	case Bamboo7:
		return "Bamboo7"
	case Bamboo8:
		return "Bamboo8"
	case Bamboo9:
		return "Bamboo9"
	case Characters1:
		return "Characters1"
	case Characters2:
		return "Characters2"
	case Characters3:
		return "Characters3"
	case Characters4:
		return "Characters4"
	case Characters5:
		return "Characters5"
	case Characters6:
		return "Characters6"
	case Characters7:
		return "Characters7"
	case Characters8:
		return "Characters8"
	case Characters9:
		return "Characters9"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	case North:
		return "North"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case White:
		return "White"
	default:
		return "Unknown"
	}
}

func checkMeld(first, second, third string) bool {
	if first == second && second == third {
		return true
	}

	iFirst, _ := strconv.Atoi(first)
	iSecond, _ := strconv.Atoi(second)
	iThird, _ := strconv.Atoi(third)

	if iFirst+1 == iSecond && iSecond == iThird-1 {
		return true
	}

	return false
}

func checkRegularWinningHands(hands string) bool {
	var number_of_meld = 0
	second_prev_pai := string(hands[0])
	first_prev_pai := string(hands[1])

	for _, _pai := range hands[2:] {
		pai := string(_pai)

		switch pai {
		case "m", "p", "s":
			continue
		}

		if second_prev_pai == "" {
			second_prev_pai = pai
			continue
		}

		if first_prev_pai == "" {
			first_prev_pai = pai
			continue
		}

		if checkMeld(second_prev_pai, first_prev_pai, pai) {
			number_of_meld += 1
			second_prev_pai = ""
			first_prev_pai = ""
		} else {
			second_prev_pai = first_prev_pai
			first_prev_pai = pai
		}
	}

	return (number_of_meld == 4)
}

func checkSevenPairs(hands string) bool {
	prevPai := string(hands[0])
	numberOfPairs := 0

	for _, _pai := range hands[1:] {
		pai := string(_pai)

		switch pai {
		case "m", "p", "s":
			continue
		}

		if prevPai == pai {
			numberOfPairs += 1
			prevPai = ""
		} else {
			prevPai = pai
		}
	}

	return numberOfPairs == 6
}

func applySuits(hand *[]PaiValue, suit rune, paiStack []rune) {
	switch suit {
	case 'm':
		for _, pai := range paiStack {
			switch pai {
			case '1':
				*hand = append(*hand, Characters1)
			case '2':
				*hand = append(*hand, Characters2)
			case '3':
				*hand = append(*hand, Characters3)
			case '4':
				*hand = append(*hand, Characters4)
			case '5':
				*hand = append(*hand, Characters5)
			case '6':
				*hand = append(*hand, Characters6)
			case '7':
				*hand = append(*hand, Characters7)
			case '8':
				*hand = append(*hand, Characters8)
			case '9':
				*hand = append(*hand, Characters9)
			}
		}
	case 'p':
		for _, pai := range paiStack {
			switch pai {
			case '1':
				*hand = append(*hand, Dots1)
			case '2':
				*hand = append(*hand, Dots2)
			case '3':
				*hand = append(*hand, Dots3)
			case '4':
				*hand = append(*hand, Dots4)
			case '5':
				*hand = append(*hand, Dots5)
			case '6':
				*hand = append(*hand, Dots6)
			case '7':
				*hand = append(*hand, Dots7)
			case '8':
				*hand = append(*hand, Dots8)
			case '9':
				*hand = append(*hand, Dots9)
			}
		}
	case 's':
		for _, pai := range paiStack {
			switch pai {
			case '1':
				*hand = append(*hand, Bamboo1)
			case '2':
				*hand = append(*hand, Bamboo2)
			case '3':
				*hand = append(*hand, Bamboo3)
			case '4':
				*hand = append(*hand, Bamboo4)
			case '5':
				*hand = append(*hand, Bamboo5)
			case '6':
				*hand = append(*hand, Bamboo6)
			case '7':
				*hand = append(*hand, Bamboo7)
			case '8':
				*hand = append(*hand, Bamboo8)
			case '9':
				*hand = append(*hand, Bamboo9)
			}
		}
	}
}

func Parse(handString string) (*[]PaiValue, error) {
	paiStack := []rune{rune(handString[0])}
	hand := []PaiValue{}

	for index, paiRune := range handString[1:] {
		switch paiRune {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			paiStack = append(paiStack, paiRune)
		case 'm', 'p', 's':
			applySuits(&hand, paiRune, paiStack)
		case 'E', 'S', 'W', 'N', 'H', 'G', 'R':
			switch paiRune {
			case 'E':
				hand = append(hand, East)
			case 'S':
				hand = append(hand, South)
			case 'W':
				hand = append(hand, West)
			case 'N':
				hand = append(hand, North)
			case 'H':
				hand = append(hand, White)
			case 'G':
				hand = append(hand, Green)
			case 'R':
				hand = append(hand, Red)
			}
		default:
			fmt.Errorf("Unexpected value %v is found at %d", string(paiRune), index)
		}
	}

	return &hand, nil
}

func JudgeTenpai(hands string) bool {
	return checkRegularWinningHands(hands) || checkSevenPairs(hands)
}

func main() {
	fmt.Println("vim-go")
}
