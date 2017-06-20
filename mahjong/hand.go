package mahjong

import "fmt"

type Hand []Pai

func (hand *Hand) groupByType() (char, dots, bamb, oner []Pai) {
	for _, pai := range *hand {
		switch pai {
		case Char1, Char2, Char3, Char4, Char5, Char6, Char7, Char8, Char9:
			char = append(char, pai)
		case Dots1, Dots2, Dots3, Dots4, Dots5, Dots6, Dots7, Dots8, Dots9:
			dots = append(dots, pai)
		case Bamb1, Bamb2, Bamb3, Bamb4, Bamb5, Bamb6, Bamb7, Bamb8, Bamb9:
			bamb = append(bamb, pai)
		case East, South, West, North, White, Green, Red:
			oner = append(oner, pai)
		}
	}

	return char, dots, bamb, oner
}

func remove(list []Pai, p Pai) []Pai {
	var result []Pai
	removed := false
	for _, e := range list {
		if e == p && !removed {
			removed = true
		} else {
			result = append(result, e)
		}
	}
	return result
}

func contain(list []Pai, p Pai) bool {
	for _, a := range list {
		if a == p {
			return true
		}
	}
	return false
}

func createCandidates(current Pai, remain []Pai, cand [][][]Pai) [][][]Pai {
	fmt.Printf("cand: %v\n", cand)
	fmt.Printf("current: %v\n", current)
	fmt.Printf("remain: %v\n", remain)
	fmt.Println("**")
	if len(remain) <= 0 {
		return cand
	}

	nextOne := current + 1
	nextTwo := current + 2

	if current.IsNumber() {
		if contain(remain, nextOne) && contain(remain, nextTwo) {
			cand[len(cand)-1] = append(cand[len(cand)-1], []Pai{current, nextOne, nextTwo})
			_remain := remove(remove(remain, nextOne), nextTwo)
			cand = createCandidates(_remain[0], _remain[1:], cand)

			base := cand[len(cand)-1]
			tmp := make([][]Pai, len(base))
			copy(tmp, base)
			cand = append(cand, tmp)
		}

		if contain(remain, nextOne) {
			cand[len(cand)-1] = append(cand[len(cand)-1], []Pai{current, nextOne})
			_remain := remove(remain, nextOne)
			cand = createCandidates(_remain[0], _remain[1:], cand)

			base := cand[len(cand)-1]
			tmp := make([][]Pai, len(base))
			copy(tmp, base)
			cand = append(cand, tmp)
		}

		if contain(remain, nextTwo) {
			cand[len(cand)-1] = append(cand[len(cand)-1], []Pai{current, nextTwo})
			_remain := remove(remain, nextTwo)
			cand = createCandidates(_remain[0], _remain[1:], cand)

			base := cand[len(cand)-1]
			tmp := make([][]Pai, len(base))
			copy(tmp, base)
			cand = append(cand, tmp)
		}
	}

	// want to remove `remove`
	if contain(remain, current) && contain(remove(remain, current), current) {
		cand[len(cand)-1] = append(cand[len(cand)-1], []Pai{current, current, current})
		_remain := remove(remove(remain, current), current)
		cand = createCandidates(_remain[0], _remain[1:], cand)

		base := cand[len(cand)-1]
		tmp := make([][]Pai, len(base))
		copy(tmp, base)
		cand = append(cand, tmp)
	}

	if contain(remain, current) {
		cand[len(cand)-1] = append(cand[len(cand)-1], []Pai{current, current})
		_remain := remove(remain, current)
		cand = createCandidates(_remain[0], _remain[1:], cand)

		base := cand[len(cand)-1]
		tmp := make([][]Pai, len(base))
		copy(tmp, base)
		cand = append(cand, tmp)
	}

	fmt.Printf("second cond: %v\n", cand)
	cand[len(cand)-1] = append(cand[len(cand)-1], []Pai{current})
	_remain := remove(remain, current)
	return createCandidates(_remain[0], _remain[1:], cand)
}

func (hand *Hand) isRegularWinningHands() bool {
	var numberOfMeld int
	var secondPrevPai Pai
	var firstPrevPai Pai

	char, dots, bamb, oner := hand.groupByType()
	fmt.Printf("hand: %v\n", hand)
	fmt.Printf("char: %v\n", char)
	fmt.Printf("dots: %v\n", dots)
	fmt.Printf("bamb: %v\n", bamb)
	fmt.Printf("oner: %v\n", oner)
	fmt.Println("---")

	_hand := *hand
	cand := [][][]Pai{[][]Pai{}}
	cand = createCandidates(_hand[0], _hand[1:], cand)

	for _, pai := range *hand {
		if secondPrevPai == Unknown {
			secondPrevPai = pai
			continue
		}

		if firstPrevPai == Unknown {
			firstPrevPai = pai
			continue
		}

		if IsMeld(secondPrevPai, firstPrevPai, pai) {
			numberOfMeld += 1
			secondPrevPai = Unknown
			firstPrevPai = Unknown
		} else {
			secondPrevPai = firstPrevPai
			firstPrevPai = pai
		}
	}

	return (numberOfMeld == 4)
}

func (hand *Hand) isSevenPairs() bool {
	var pairs []Pai
	var prevPai Pai

	appendIfMissing := func(pairs []Pai, pai Pai) []Pai {
		for _, p := range pairs {
			if p == pai {
				return pairs
			}
		}

		return append(pairs, pai)
	}

	for _, pai := range *hand {
		if prevPai == pai {
			pairs = appendIfMissing(pairs, pai)
			prevPai = Unknown
		} else {
			prevPai = pai
		}
	}

	return len(pairs) == 6
}

func (hand *Hand) isThirteenOrphans() bool {
	var tmpHand Hand

	appendIfMissing := func(hand Hand, pai Pai) Hand {
		for _, p := range hand {
			if p == pai {
				return hand
			}
		}

		return append(hand, pai)
	}

	for _, pai := range *hand {
		if pai.IsOrphan() {
			tmpHand = appendIfMissing(tmpHand, pai)
		} else {
			return false
		}
	}

	return len(tmpHand) == 13 || len(tmpHand) == 12
}

func (hand *Hand) IsTenpai() bool {
	return hand.isRegularWinningHands() ||
		hand.isSevenPairs() ||
		hand.isThirteenOrphans()
}
