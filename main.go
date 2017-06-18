package main

import (
	"fmt"

	"github.com/tanaka51/tenpai-wakaru-man/mahjong"
)

func checkRegularWinningHands(hand *mahjong.Hand) bool {
	var numberOfMeld int
	var secondPrevPai mahjong.Pai
	var firstPrevPai mahjong.Pai

	for _, pai := range *hand {
		if secondPrevPai == mahjong.Unknown {
			secondPrevPai = pai
			continue
		}

		if firstPrevPai == mahjong.Unknown {
			firstPrevPai = pai
			continue
		}

		if mahjong.IsMeld(secondPrevPai, firstPrevPai, pai) {
			numberOfMeld += 1
			secondPrevPai = mahjong.Unknown
			firstPrevPai = mahjong.Unknown
		} else {
			secondPrevPai = firstPrevPai
			firstPrevPai = pai
		}
	}

	return (numberOfMeld == 4)
}

func checkSevenPairs(hand *mahjong.Hand) bool {
	var pairs []mahjong.Pai
	var prevPai mahjong.Pai

	appendIfMissing := func(pairs []mahjong.Pai, pai mahjong.Pai) []mahjong.Pai {
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
			prevPai = mahjong.Unknown
		} else {
			prevPai = pai
		}
	}

	return len(pairs) == 6
}

func JudgeTenpai(hands string) bool {
	hand, _ := mahjong.Parse(hands)
	return checkRegularWinningHands(hand) || checkSevenPairs(hand)
}

func main() {
	fmt.Println("vim-go")
}
