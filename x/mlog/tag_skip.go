package mlog

import (
	"strings"
)

var skips = map[string]bool{}

//IsSkip : Check tag is skip
func IsSkip(tag string) bool {
	return skips[safeTagName(tag)]
}

//SkipTag : skip
func SkipTag(tag string) {
	skips[safeTagName(tag)] = true
}

func safeTagName(tag string) string {
	return strings.ToLower(tag)
}
