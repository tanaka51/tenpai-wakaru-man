package pai

import (
	"reflect"
	"testing"
)

var tests = []struct {
	handString string
	expected   []Value
}{
	{"123456789mESWN", []Value{Char1, Char2, Char3, Char4, Char5, Char6, Char7, Char8, Char9, East, South, West, North}},
	{"123m123p123sRGH1m", []Value{Char1, Char1, Char2, Char3, Dots1, Dots2, Dots3, Bamb1, Bamb2, Bamb3, White, Green, Red}},
	{"123123mRRRGGGE", []Value{Char1, Char1, Char2, Char2, Char3, Char3, East, Green, Green, Green, Red, Red, Red}},
	{"S9pH8m7sW654pN32m1s", []Value{Char2, Char3, Char8, Dots4, Dots5, Dots6, Dots9, Bamb1, Bamb7, West, North, White}},
}

func TestParse(t *testing.T) {
	for _, test := range tests {
		hand, _ := Parse(test.handString)
		if !reflect.DeepEqual(*hand, test.expected) {
			t.Errorf("%s is not parsed correctly. expected %v but %v", test.handString, test.expected, hand)
		}
	}
}
