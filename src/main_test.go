package main

import (
	"testing"
)

var testcases = []struct {
	hands    string
	expected bool
}{
	{"123456789mEEEW", true},
	{"13579m2468pRGWES", false},
}

func TestJudgeTenpai(t *testing.T) {
	for _, testcase := range testcases {
		result := JudgeTenpai(testcase.hands)
		if result != testcase.expected {
			t.Errorf("%s must be %v", testcase.hands, testcase.expected)
		}
	}
}
