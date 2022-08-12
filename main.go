package main

import "fmt"

func main() {
	number := 5
	fmt.Println(number)
	change(&number, 100)
	fmt.Println(number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("didalam function:", *old)
}
