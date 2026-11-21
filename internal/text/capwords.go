package text

import "strings"

func CapitalizeWords(content string) string {
	words := strings.Fields(content) // Split the content into words based on whitespace

	var upperWords []string

	for i := 0; i < len(words); i++ {
		// letters := strings.Split(words[i], "") // Split each word into individual letters
		// if len(letters) > 0 {
		// 	letters[0] = strings.ToUpper(letters[0]) // Capitalize the first letter of each word
		// 	upperWords = append(upperWords, strings.Join(letters, ""))      // Join the letters back into a word
			letters := []rune(words[i]) // Convert the word into a slice of runes to handle Unicode characters
		if len(letters) > 0 {
			letters[0] = []rune(strings.ToUpper(string(letters[0])))[0] // Capitalize the first letter of each word
			upperWords = append(upperWords, string(letters)) // Join the letters back into a word
		} else {
			upperWords = append(upperWords, words[i]) // If the word is empty, just append it as is
		}
	}
	return strings.Join(upperWords, " ")
}
