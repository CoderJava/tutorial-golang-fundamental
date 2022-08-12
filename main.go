package main

import "fmt"

func main() {
	numberA := 5
	numberB := &numberA // reference
	numberC := *numberB // dereference

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
	fmt.Println(numberC)
	fmt.Println("====================")

	*numberB = 10
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
	fmt.Println(numberC)
}
