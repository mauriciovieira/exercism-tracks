// Package raindrops exports the Convert function for generating raindrops
package raindrops

import (
	"bytes"
	"sort"
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

	var keys []int
	for k := range dictionary {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, index := range keys {
		if number%index == 0 {
			buffer.WriteString(dictionary[index])
		}
	}
	if buffer.Len() == 0 {
		return strconv.Itoa(number)
	}
	return buffer.String()
}
