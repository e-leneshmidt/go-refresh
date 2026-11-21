package text

import "strings"

func CountWords(content string) int {
	
	text := strings.Fields(content) //Fields function splits the string into a slice of words based on whitespace//

	return len(text) //len function returns the number of elements in the slice, which is the count of words in the content//
}
