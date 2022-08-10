package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Golang the best language"
	for index, letter := range sentence {
		if index%2 == 0 {
			strLetter := string(letter)
			fmt.Println("index:", index, "letter:", strLetter)
		}
	}

	fmt.Println()
	vocalCharacters := []string{
		"a",
		"i",
		"u",
		"e",
		"o",
	}
	for _, letter := range sentence {
		strLetter := string(letter)
		strLetterLowerCase := strings.ToLower(strLetter)
		for _, letterVocal := range vocalCharacters {
			if (strLetterLowerCase == letterVocal) {
				fmt.Println(strLetter)
			}
		}
	}
}
