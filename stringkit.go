package fiberparser

import (
	"strings"
)

// stringkitAfter returns a substring after a string. Or empty if it can't find it.
func stringkitAfter(str string, subStrAfter string) string {
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

// stringkitBetween returns a substring between two strings. Or empty if it can't find it.
func stringkitBetween(str string, subStrBefore string, subStrAfter string) string {
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
