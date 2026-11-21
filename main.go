package main

import (
	"fmt"
	"go-refresh/internal/file"
	"go-refresh/internal/text"
)

func main() {
	content, err := file.ReadFile("internal/testdata/input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(content)

	fmt.Println(text.CountWords(content))
	
	err = file.WriteFile("internal/testdata/output.txt", "Hello World")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("File written successfully!")

	allUpper := text.ToUpperWords(content)
	fmt.Println(allUpper)

	allLower := text.ToLowerWords(content)
	fmt.Println(allLower)

	capitalized := text.CapitalizeWords(content)
	fmt.Println(capitalized)

	reversed := text.ReverseWords(content)
	fmt.Println(reversed)

}

