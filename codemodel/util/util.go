package util

import (
	"unicode"
)

// Make first letter Cap
func Capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	s = string(a)
	return s
}
