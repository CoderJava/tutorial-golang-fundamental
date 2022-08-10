package main

import "fmt"

func main() {
	// if else
	/* age := 28
	if age > 10 {
		fmt.Println("Boleh bermain game")
	} else {
		fmt.Println("Tidak boleh bermain game")
	} */

	// else if
	score := 80
	var grade string
	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else if score <= 85 {
		grade = "B"
	} else {
		grade = "A"
	}
	fmt.Println("grade:", grade)
}
