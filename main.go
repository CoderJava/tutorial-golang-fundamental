package main

import (
	"fmt"
	"golang-fundamental/calculation"
)

func main() {
	resultAdd := calculation.Add(5, 8)
	fmt.Println("result add:", resultAdd)

	resultMultiply := calculation.Multiply(4, 4)
	fmt.Println("result multiply:", resultMultiply)
}
