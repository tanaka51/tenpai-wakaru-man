package main

import (
	"fmt"
	"strconv"
)

func JudgeTenpai(hands string) bool {
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

		second_prev_pai_int, _ := strconv.Atoi(second_prev_pai)
		first_prev_pai_int, _ := strconv.Atoi(first_prev_pai)
		pai_int, _ := strconv.Atoi(pai)

		if ((second_prev_pai_int + 1) == first_prev_pai_int) && (first_prev_pai_int == (pai_int - 1)) {
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

func main() {
	fmt.Println("vim-go")
}
