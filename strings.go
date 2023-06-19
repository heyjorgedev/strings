package strings

import (
	"strings"
	"unicode/utf8"
)

func Lower(str string) string {
	return strings.ToLower(str)
}

func Upper(str string) string {
	return strings.ToUpper(str)
}

func After(original, after string) string {
	str := strings.Split(original, after)
	if len(str) > 1 {
		return str[1]
	}

	return ""
}

func Before(original, before string) string {
	str := strings.Split(original, before)
	return str[0]
}

func Contains(str, needle string) bool {
	return strings.Contains(str, needle)
}

func ContainsSome(str string, needles []string) bool {
	for _, needle := range needles {
		if Contains(str, needle) {
			return true
		}
	}

	return false
}

func ContainsAll(str string, needles []string) bool {
	for _, needle := range needles {
		if !Contains(str, needle) {
			return false
		}
	}

	return true
}

func Length(str string) int {
	return utf8.RuneCountInString(str)
}

func Reverse(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
