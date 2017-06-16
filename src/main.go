package main

import (
	"fmt"
	"strconv"
)

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

func JudgeTenpai(hands string) bool {
	return checkRegularWinningHands(hands) || checkSevenPairs(hands)
}

func main() {
	fmt.Println("vim-go")
}
