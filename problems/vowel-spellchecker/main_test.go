package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spellchecker(t *testing.T) {
	type testCase struct {
		Wordlist []string
		Queries  []string
		Answer   []string
	}

	testCases := []testCase{
		{
			Wordlist: []string{"KiTe", "kite", "hare", "Hare"},
			Queries:  []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			Answer:   []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Answer, spellchecker(tc.Wordlist, tc.Queries))
	}
}
