package main

import "fmt"

func main() {
	/* myMap := map[string]int{}

	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["JavaScript"] = 8
	myMap["go"] = 10 */

	// fmt.Println(myMap)
	// fmt.Println(myMap["go"])

	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"JavaScript": "hype",
	}
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println("==================")

	delete(myMap, "ruby")
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	value, isAvailable := myMap["python"]
	fmt.Println(value)
	fmt.Println(isAvailable)
}
