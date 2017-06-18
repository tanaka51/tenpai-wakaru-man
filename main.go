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
	var numberOfPairs int
	var prevPai mahjong.Pai

	for _, pai := range *hand {
		if prevPai == pai {
			numberOfPairs += 1
			prevPai = mahjong.Unknown
		} else {
			prevPai = pai
		}
	}

	return numberOfPairs == 6
}

func JudgeTenpai(hands string) bool {
	hand, _ := mahjong.Parse(hands)
	return checkRegularWinningHands(hand) || checkSevenPairs(hand)
}

func main() {
	fmt.Println("vim-go")
}
