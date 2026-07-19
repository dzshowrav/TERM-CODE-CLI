package util

// Truncate truncates a string to maxLen runes with a middle ellipsis.
func Truncate(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen < 5 {
		return string(runes[:maxLen])
	}
	left := maxLen/2 - 1
	right := maxLen/2 - 2
	if left+right+3 > maxLen {
		right = maxLen - left - 3
	}
	return string(runes[:left]) + "..." + string(runes[len(runes)-right:])
}
