// Package hamming defines Hamming Distance function
package hamming

import (
	"errors"
)

// Distance calculcates the hamming distance between two sequences
func Distance(a, b string) (d int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("strings have different size")
	}

	for i := range a {
		if a[i] != b[i] {
			d += 1
		}
	}

	return
}
