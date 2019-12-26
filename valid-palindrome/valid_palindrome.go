package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = reg.ReplaceAllString(s, "")

	if len(s) <= 1 {
		return true
	}

	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
