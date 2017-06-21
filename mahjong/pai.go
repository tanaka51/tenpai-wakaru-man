package mahjong

import (
	"fmt"
	"sort"
)

type Pai int
type Suit int

const (
	Unknown Pai = iota
	Char1
	Char2
	Char3
	Char4
	Char5
	Char6
	Char7
	Char8
	Char9
	Dots1
	Dots2
	Dots3
	Dots4
	Dots5
	Dots6
	Dots7
	Dots8
	Dots9
	Bamb1
	Bamb2
	Bamb3
	Bamb4
	Bamb5
	Bamb6
	Bamb7
	Bamb8
	Bamb9
	East
	South
	West
	North
	White
	Green
	Red
)

const (
	UnknownSuit Suit = iota + 100
	Char
	Dots
	Bamb
)

func (p Pai) String() string {
	switch p {
	case Char1:
		return "Char1"
	case Char2:
		return "Char2"
	case Char3:
		return "Char3"
	case Char4:
		return "Char4"
	case Char5:
		return "Char5"
	case Char6:
		return "Char6"
	case Char7:
		return "Char7"
	case Char8:
		return "Char8"
	case Char9:
		return "Char9"
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
	case Bamb1:
		return "Bamb1"
	case Bamb2:
		return "Bamb2"
	case Bamb3:
		return "Bamb3"
	case Bamb4:
		return "Bamb4"
	case Bamb5:
		return "Bamb5"
	case Bamb6:
		return "Bamb6"
	case Bamb7:
		return "Bamb7"
	case Bamb8:
		return "Bamb8"
	case Bamb9:
		return "Bamb9"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	case North:
		return "North"
	case White:
		return "White"
	case Green:
		return "Green"
	case Red:
		return "Red"
	default:
		return "Unknown"
	}
}

func (s Suit) String() string {
	switch s {
	case Char:
		return "Char"
	case Dots:
		return "Dots"
	case Bamb:
		return "Bamb"
	default:
		return "Unknown"
	}
}

func (p Pai) Suit() Suit {
	switch p {
	case Char1, Char2, Char3, Char4, Char5, Char6, Char7, Char8, Char9:
		return Char
	case Dots1, Dots2, Dots3, Dots4, Dots5, Dots6, Dots7, Dots8, Dots9:
		return Dots
	case Bamb1, Bamb2, Bamb3, Bamb4, Bamb5, Bamb6, Bamb7, Bamb8, Bamb9:
		return Bamb
	default:
		return UnknownSuit
	}
}

func (p Pai) IsNumber() bool {
	switch p {
	case Char1, Char2, Char3, Char4, Char5, Char6, Char7, Char8, Char9:
		return true
	case Dots1, Dots2, Dots3, Dots4, Dots5, Dots6, Dots7, Dots8, Dots9:
		return true
	case Bamb1, Bamb2, Bamb3, Bamb4, Bamb5, Bamb6, Bamb7, Bamb8, Bamb9:
		return true
	default:
		return false
	}
}

func (p Pai) IsOrphan() bool {
	switch p {
	case Char1, Char9, Dots1, Dots9, Bamb1, Bamb9:
		return true
	case East, South, West, North:
		return true
	case White, Green, Red:
		return true
	default:
		return false
	}
}

func applySuits(hand *Hand, suit rune, paiStack []rune) {
	switch suit {
	case 'm':
		for _, pai := range paiStack {
			switch pai {
			case '1':
				*hand = append(*hand, Char1)
			case '2':
				*hand = append(*hand, Char2)
			case '3':
				*hand = append(*hand, Char3)
			case '4':
				*hand = append(*hand, Char4)
			case '5':
				*hand = append(*hand, Char5)
			case '6':
				*hand = append(*hand, Char6)
			case '7':
				*hand = append(*hand, Char7)
			case '8':
				*hand = append(*hand, Char8)
			case '9':
				*hand = append(*hand, Char9)
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
				*hand = append(*hand, Bamb1)
			case '2':
				*hand = append(*hand, Bamb2)
			case '3':
				*hand = append(*hand, Bamb3)
			case '4':
				*hand = append(*hand, Bamb4)
			case '5':
				*hand = append(*hand, Bamb5)
			case '6':
				*hand = append(*hand, Bamb6)
			case '7':
				*hand = append(*hand, Bamb7)
			case '8':
				*hand = append(*hand, Bamb8)
			case '9':
				*hand = append(*hand, Bamb9)
			}
		}
	}
}

func Parse(handString string) (*Hand, error) {
	paiStack := []rune{}
	hand := Hand{}

	for index, paiRune := range handString {
		switch paiRune {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			paiStack = append(paiStack, paiRune)
		case 'm', 'p', 's':
			applySuits(&hand, paiRune, paiStack)
			paiStack = []rune{}
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

	sort.Slice(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})

	return &hand, nil
}

func IsMeld(first, second, third Pai) bool {
	if first == second && second == third {
		return true
	}

	if !(first.IsNumber() && second.IsNumber() && third.IsNumber()) {
		return false
	}

	if (first+2) == (second+1) && (second+1) == (third) {
		return true
	}

	if (first) == (second+1) && (second+1) == (third+2) {
		return true
	}

	return false
}
