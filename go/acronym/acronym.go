// Package acronym is a package for generating acronyms.
package acronym

import (
	"strings"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	words := strings.Fields(strings.Title(strings.Replace(s, "-", " ", -1)))
	letters := make([]byte, len(words))

	for idx, word := range words {
		letters[idx] = word[0]
	}

	return string(letters[:])
}
