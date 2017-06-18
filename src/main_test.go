package main

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	hands    string
	expected bool
}{
	{"123456789m123p1s", true},
	{"111222333m111p1s", true},
	{"EEESSSWWWNNNW", true},
	{"13579m2468p1357s", false},
	{"13579m2468pESWN", false},
	{"1133557799mEEN", true},
}

func TestJudgeTenpai(t *testing.T) {
	for _, testcase := range testcases {
		result := JudgeTenpai(testcase.hands)
		if result != testcase.expected {
			t.Errorf("%s must be %v but %v", testcase.hands, testcase.expected, result)
		}
	}
}

var tests2 = []struct {
	handString string
	expected   []PaiValue
}{
	{"123456789mESWN", []PaiValue{Characters1, Characters2, Characters3, Characters4, Characters5, Characters6, Characters7, Characters8, Characters9, East, South, West, North}},
}

func TestParse(t *testing.T) {
	for _, test := range tests2 {
		hand, _ := Parse(test.handString)
		if !reflect.DeepEqual(*hand, test.expected) {
			t.Errorf("%s is not parsed correctly. result: %v", test.handString, hand)
		}
	}
}
