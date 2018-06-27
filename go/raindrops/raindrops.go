// Package raindrops exports the Convert function for generating raindrops
package raindrops

import (
	"bytes"
	"strconv"
)

var dictionary = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert is function that converts numbers into strings in raindrop language
func Convert(number int) string {
	var buffer bytes.Buffer
	for index, drop := range dictionary {
		if number%index == 0 {
			buffer.WriteString(drop)
		}
	}
	if buffer.Len() == 0 {
		return strconv.Itoa(number)
	}
	return buffer.String()
}