package main

import "fmt"

func main() {
	result := add(10, 20)
	fmt.Println(result)
}

func add(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}