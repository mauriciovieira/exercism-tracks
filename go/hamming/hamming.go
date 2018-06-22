// Package hamming defines Hamming Distance function
package hamming

import (
	"errors"
	"strings"
)

// Distance calculcates the hamming distance between two sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings have different size")
	}

	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")

	distance := 0

	for i := range aArr {
		if aArr[i] != bArr[i] {
			distance += 1
		}
	}

	return distance, nil
}
