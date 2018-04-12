// Package bob implements a simple chatbot
package bob

import "regexp"

// Hey bob, answer me
func Hey(remark string) string {
	toBeRemoved := regexp.MustCompile(`[\s!,]`)
	subject := toBeRemoved.ReplaceAllLiteralString(remark, "")

	isAQuestion := regexp.MustCompile(`\?$`)
	isShouting := regexp.MustCompile(`^([^a-z])+$`)
	isAnxious := regexp.MustCompile(`^[A-Z]+\?$`)
	onlyNumbers := regexp.MustCompile(`^[0-9]+$`)

	var answer string

	switch {
	case onlyNumbers.MatchString(subject):
		answer = "Whatever."
	case isAnxious.MatchString(subject):
		answer = "Calm down, I know what I'm doing!"
	case isAQuestion.MatchString(subject):
		answer = "Sure."
	case isShouting.MatchString(subject):
		answer = "Whoa, chill out!"
	case subject == "":
		answer = "Fine. Be that way!"
	default:
		answer = "Whatever."
	}
	return answer
}
