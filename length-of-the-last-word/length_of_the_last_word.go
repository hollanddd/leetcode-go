package lengthlastword

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(strings.TrimSpace(s)) == 0 {
		return 0
	}

	parts := strings.Split(strings.TrimSpace(s), " ")

	if len(parts) == 1 {
		return len(parts[0])
	}

	last := len(parts) - 1

	return len(parts[last])
}
