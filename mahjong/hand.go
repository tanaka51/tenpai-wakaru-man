package mahjong

type Hand []Pai

func (hand *Hand) isRegularWinningHands() bool {
	var numberOfMeld int
	var secondPrevPai Pai
	var firstPrevPai Pai

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
