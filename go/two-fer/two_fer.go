// Package twofer defines ShareWith

package twofer

import (
	"fmt"
)

// ShareWith returns a phrase
func ShareWith(name string) string {
	person := name

	if name == "" {
		person = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", person)
}
