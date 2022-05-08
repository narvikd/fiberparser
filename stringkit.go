package fiberparser

import (
	"strings"
)

// After returns a substring after a string. Or empty if it can't find it.
func After(str string, subStrAfter string) string {
	pos := strings.LastIndex(str, subStrAfter)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(subStrAfter)
	if adjustedPos >= len(str) {
		return ""
	}
	return str[adjustedPos:]
}

// Between returns a substring between two strings. Or empty if it can't find it.
func Between(str string, subStrBefore string, subStrAfter string) string {
	posFirst := strings.Index(str, subStrBefore)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(str, subStrAfter)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(subStrBefore)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return str[posFirstAdjusted:posLast]
}