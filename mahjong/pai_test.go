package mahjong

import (
	"reflect"
	"testing"
)

var tests = []struct {
	handString string
	expected   Hand
}{
	{"123456789mESWN", Hand{Char1, Char2, Char3, Char4, Char5, Char6, Char7, Char8, Char9, East, South, West, North}},
	{"123m123p123sRGH1m", Hand{Char1, Char1, Char2, Char3, Dots1, Dots2, Dots3, Bamb1, Bamb2, Bamb3, White, Green, Red}},
	{"123123mRRRGGGE", Hand{Char1, Char1, Char2, Char2, Char3, Char3, East, Green, Green, Green, Red, Red, Red}},
	{"S9pH8m7sW654pN32m1s", Hand{Char2, Char3, Char8, Dots4, Dots5, Dots6, Dots9, Bamb1, Bamb7, South, West, North, White}},
	{"EEESSSWWWNNNH", Hand{East, East, East, South, South, South, West, West, West, North, North, North, White}},
}

func TestParse(t *testing.T) {
	for _, test := range tests {
		hand, _ := Parse(test.handString)
		if !reflect.DeepEqual(*hand, test.expected) {
			t.Errorf("%s is not parsed correctly. expected %v but %v", test.handString, test.expected, hand)
		}
	}
}

var testsMeld = []struct {
	first    Pai
	second   Pai
	third    Pai
	expected bool
}{
	{Char1, Char2, Char3, true},
	{Char1, Char1, Char1, true},
	{Char9, Char8, Char7, true},
	{East, East, East, true},
	{White, Green, Red, false},
	{White, White, White, true},
	{Dots2, Dots2, Dots2, true},
	{Bamb6, Bamb6, Bamb6, true},
	{Char2, Dots2, Bamb2, false},
}

func TestIsMeld(t *testing.T) {
	for _, test := range testsMeld {
		if result := IsMeld(test.first, test.second, test.third); result != test.expected {
			t.Errorf("%s %s %s should be %v but %v", test.first, test.second, test.third, test.expected, result)
		}
	}
}
