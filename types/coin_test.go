package types

import (
	"testing"
)

func TestValidDenoms(t *testing.T) {
	shr := NewCoin("SHR", 1)
	shrp := NewCoin("SHRP", 1)
	invalid := NewCoin("Invalid", 1)

	table := []struct {
		input    Coins
		expected bool
	}{
		{Coins([]Coin{shr, shrp}), true},
		{Coins([]Coin{shr}), false},
		{Coins([]Coin{shrp}), false},
		{Coins([]Coin{shr, shr}), false},
		{Coins([]Coin{shr, invalid}), false},
		{Coins([]Coin{shr, shr, shrp}), false},
	}

	for _, tc := range table {
		ret := tc.input.HasValidDenoms()
		if ret != tc.expected {
			t.Logf("%s HasValidDenoms should return %t but %t returned.", tc.input, tc.expected, ret)
		}
	}
}

func TestCoinsBothDenoms(t *testing.T) {

	coin1 := NewCoin("SHR", 1)

	coin := NewCoin("SHR", 1)
	coins3 := coin1.Plus(coin)

	t.Logf("%s\n", coins3)

	coins4 := coins3.Minus(coin)
	t.Logf("%s\n", coins4)

	coin = NewCoin("SHRP", 2)
	coins3 = coin1.Plus(coin)
	coins4 = coins4.Minus(coin)
	t.Logf("%s\n", coins3)
	t.Logf("%s\n", coins4)
}
