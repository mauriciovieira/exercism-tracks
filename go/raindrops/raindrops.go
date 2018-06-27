// Package raindrops exports the Convert function for generating raindrops
package raindrops

import (
	"strconv"
)

// waterdropGenerator is used internally to expose various attributes
type waterdropGenerator struct {
	drop int
}

// isPling returns true if the drop is divisible by 3
func (r *waterdropGenerator) isPling() bool {
	return r.drop%3 == 0
}

// isPlang returns true if the drop is divisible by 5
func (r *waterdropGenerator) isPlang() bool {
	return r.drop%5 == 0
}

// isPlang returns true if the drop is divisible by 8
func (r *waterdropGenerator) isPlong() bool {
	return r.drop%7 == 0
}

// Convert will take a string and return the snarky response of a teenager named Bob
func Convert(drop int) string {
	waterdrop := &waterdropGenerator{drop}
	remarks := ""

	if waterdrop.isPling() {
		remarks += "Pling"
	}

	if waterdrop.isPlang() {
		remarks += "Plang"
	}

	if waterdrop.isPlong() {
		remarks += "Plong"
	}

	if remarks == "" {
		return strconv.Itoa(drop)
	}
	return remarks
}
