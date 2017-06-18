package main

import (
	"testing"
)

var testcases = []struct {
	hands    string
	expected bool
}{
	{"123456789m123p1s", true},
	{"111222333m111p1s", true},
	{"EEESSSWWWNNNH", true},
	{"13579m2468p1357s", false},
	{"13579m2468pESWN", false},
	{"1133557799mEEN", true},
	{"1111557799mEEN", false},
}

func TestJudgeTenpai(t *testing.T) {
	for _, testcase := range testcases {
		result := JudgeTenpai(testcase.hands)
		if result != testcase.expected {
			t.Errorf("%s must be %v but %v", testcase.hands, testcase.expected, result)
		}
	}
}
