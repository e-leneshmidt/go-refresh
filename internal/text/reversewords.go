package text

import "strings"

func SplitWords(content string) []string {
	words := strings.Fields(content)
	return words
}

func ReverseWords(content string) string {
	words := SplitWords(content)
	if len(words) == 0 {
		return ""
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
