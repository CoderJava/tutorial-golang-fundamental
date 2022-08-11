package main

import "fmt"

func main() {
	/* var languages [5]string
	languages[0] = "Go"
	languages[1] = "Ruby"
	languages[2] = "JavaScript"
	languages[3] = "C#"
	languages[4] = "Python" */

	/* languages := [5]string{
		"Go",
		"Ruby",
		"JavaScript",
		"C#",
		"Python",
	} */

	languages := [...]string{
		"Ruby",
		"Python",
		"JavaScript",
		"Go",
		"C#",
		"Kotlin",
	}

	fmt.Println(languages)
	length := len(languages)
	fmt.Println(length)
}
