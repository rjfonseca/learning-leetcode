package main

import (
	"strings"
)

func main() {
}

func spellchecker(wordlist []string, queries []string) []string {
	answer := make([]string, len(queries))
	// We will create 3 different indexes
	// 1- Original word
	// 2- Lower cased
	// 3- Masked vowels
	index := make(map[string]int, len(wordlist))
	lowerIdx := make(map[string]int, len(wordlist))
	maskedIdx := make(map[string]int, len(wordlist))

	// Build de indexes
	for i, word := range wordlist {
		index[word] = i
		lower := strings.ToLower(word)
		if _, ok := lowerIdx[lower]; !ok {
			lowerIdx[lower] = i
		}
		masked := maskVowels(lower)
		if _, ok := maskedIdx[masked]; !ok {
			maskedIdx[masked] = i
		}
	}

	// Search words
	for i, query := range queries {
		// Search for the original word
		if j, ok := index[query]; ok {
			answer[i] = wordlist[j]
			continue
		}

		// Search lower case
		lower := strings.ToLower(query)
		if j, ok := lowerIdx[lower]; ok {
			answer[i] = wordlist[j]
			continue
		}

		// Search masked vowels.
		masked := maskVowels(lower)
		if j, ok := maskedIdx[masked]; ok {
			answer[i] = wordlist[j]
		}

	}

	return answer
}

// maskVowels replaces any vowels with '*'.
// It assumes the string s is already lower case.
func maskVowels(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			r = '*'
		}
		sb.WriteRune(r)
	}

	return sb.String()
}
