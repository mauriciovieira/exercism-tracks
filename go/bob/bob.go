// Package bob exports the Hey function that will respond to a string as a snarky teenager would
package bob

import "strings"

// remarkAnalyzer is used internally to expose various attributes
type remarkAnalyzer struct {
	remark string
}

// isYelling returns true if a remark is in all caps and alpha characters exist in the string
func (r *remarkAnalyzer) isYelling() bool {
	return strings.ToUpper(r.remark) == r.remark && strings.ToLower(r.remark) != r.remark
}

// isQuestion returns true if a remark ends with a question mark
func (r *remarkAnalyzer) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

// isSilent returns true if the remark is empty
func (r *remarkAnalyzer) isSilent() bool {
	return r.remark == ""
}

// Hey will take a string and return the snarky response of a teenager named Bob
func Hey(remark string) string {
	analyzer := &remarkAnalyzer{strings.TrimSpace(remark)}
	switch {
	case analyzer.isQuestion() && analyzer.isYelling():
		return "Calm down, I know what I'm doing!"
	case analyzer.isQuestion():
		return "Sure."
	case analyzer.isYelling():
		return "Whoa, chill out!"
	case analyzer.isSilent():
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}