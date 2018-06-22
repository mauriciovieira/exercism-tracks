// Package hamming defines Hamming Distance function
package hamming

import (
	"errors"
)

// Distance calculcates the hamming distance between two sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings have different size")
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance, nil
}
