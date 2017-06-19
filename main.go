package main

import (
	"fmt"

	"github.com/tanaka51/tenpai-wakaru-man/mahjong"
)

func JudgeTenpai(hands string) bool {
	hand, _ := mahjong.Parse(hands)
	return hand.IsTenpai()
}

func main() {
	fmt.Println("vim-go")
}
