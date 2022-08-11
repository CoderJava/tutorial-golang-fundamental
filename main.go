package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{
		10,
		5,
		8,
		9,
		7,
	}
	total := sum(scores)
	fmt.Println("total:", total)

	hasilPenjumlahan, errPenjumlahan := calculate(10, 2, "+")
	if errPenjumlahan != nil {
		fmt.Println("error penjumlahan:", errPenjumlahan)
	} else {
		fmt.Println("hasil penjumlahan:", hasilPenjumlahan)
	}

	hasilPengurangan, errPengurangan := calculate(10, 2, "-")
	if errPengurangan != nil {
		fmt.Println("error pengurangan:", errPengurangan)
	} else {
		fmt.Println("hasil pengurangan:", hasilPengurangan)
	}

	hasilPerkalian, errPerkalian := calculate(10, 2, "*")
	if errPerkalian != nil {
		fmt.Println("error perkalian:", errPerkalian)
	} else {
		fmt.Println("hasil perkalian:", hasilPerkalian)
	}

	hasilPembagian, errPembagian := calculate(10, 2, "/")
	if errPembagian != nil {
		fmt.Println("error pembagian:", errPembagian)
	} else {
		fmt.Println("hasil pembagian:", hasilPembagian)
	}

	hasilModulo, errModulo := calculate(10, 2, "%")
	if errModulo != nil {
		fmt.Println("error modulo:", errModulo)
	} else {
		fmt.Println("hasil modulo:", hasilModulo)
	}

}

func sum(values []int) int {
	var output int
	for _, value := range values {
		output += value
	}
	return output
}

func calculate(numberOne int, numberTwo int, operation string) (output int, err error) {
	switch operation {
	case "+":
		output = numberOne + numberTwo
	case "-":
		output = numberOne - numberTwo
	case "*":
		output = numberOne * numberTwo
	case "/":
		output = numberOne / numberTwo
	default:
		errorMessage := fmt.Sprintf("argument %s operation not identified", operation)
		err = errors.New(errorMessage)
	}
	return
}
