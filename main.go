package main

import "fmt"

func main() {
	/* number := 2
	switch number {
	case 1:
		fmt.Println("S")
		fmt.Println("A")
		fmt.Println("T")
		fmt.Println("U")
	case 2:
		fmt.Println("D")
		fmt.Println("U")
		fmt.Println("A")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("DEFAULT")
	} */

	score := 88
	var grade string
	switch {
	case score <= 50:
		grade = "E"
	case score <= 60:
		grade = "D"
	case score <= 70:
		grade = "C"
	case score <= 80:
		grade = "B"
	default:
		grade = "A"
	}
	fmt.Println(grade)
}
