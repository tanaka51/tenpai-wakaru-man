package mahjong

import "testing"

var regularWiningHandTests = []struct {
	hand     *Hand
	expected bool
}{
	{
		&Hand{
			Char1, Char1, Char2, Char2, Char3, Char3, Char4, Char5, Char6, Char7, Char8, Char9, Char9,
		}, true,
	},
	{
		&Hand{
			Char1, Bamb1, Dots1, Char5, Bamb5, Dots5, Char9, Bamb9, Dots9, North, Red, Red, White,
		}, false,
	},
	{
		&Hand{
			Char1, Char2, Char3, Char7, Char8, Dots7, Dots7, Dots7, Bamb5, Bamb6, Bamb7, North, North,
		}, true,
	},
	{
		&Hand{
			Char1, Char2, Char3, Char4, Char5, Char6, Char7, Char8, Char9, Dots1, Dots2, Dots3, North,
		}, true,
	},
}

func TestIsRegularWinningHands(t *testing.T) {
	for _, test := range regularWiningHandTests {
		if result := test.hand.isRegularWinningHands(); result != test.expected {
			t.Errorf("expected %v but got %v. hands: %v", test.expected, result, test.hand)
		}
	}
}

var sevenPairsHandTests = []struct {
	hand     *Hand
	expected bool
}{
	{
		&Hand{
			Char1, Char1, Char2, Char2, Char3, Char3, Char4, Char4, Char5, Char5, Bamb1, Bamb1, North,
		}, true,
	},
	{
		&Hand{
			Char1, Char1, Char1, Char1, Char3, Char3, Char4, Char4, Char5, Char5, Bamb1, Bamb1, North,
		}, false,
	},
}

func TestIsSevenPairs(t *testing.T) {
	for _, test := range sevenPairsHandTests {
		if result := test.hand.isSevenPairs(); result != test.expected {
			t.Errorf("expected %v but got %v. hands: %v", test.expected, result, test.hand)
		}
	}
}

var thirteenOrphansTests = []struct {
	hand     *Hand
	expected bool
}{
	{
		&Hand{
			Char1, Char9, Dots1, Dots9, Bamb1, Bamb9, East, South, West, North, White, Green, Red,
		}, true,
	},
	{
		&Hand{
			Char1, Char9, Dots1, Dots9, Bamb1, Bamb9, East, South, West, North, White, Green, Green,
		}, true,
	},
	{
		&Hand{
			Char1, Char2, Char9, Dots1, Dots9, Bamb1, Bamb9, East, South, West, North, White, Green,
		}, false,
	},
}

func TestIsThirteenOrphans(t *testing.T) {
	for _, test := range thirteenOrphansTests {
		if result := test.hand.isThirteenOrphans(); result != test.expected {
			t.Errorf("expected %v but got %v. hands: %v", test.expected, result, test.hand)
		}
	}
}
