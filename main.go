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
		"ruby": "is beautiful",
		"go": "is super fast",
	}
	fmt.Println(myMap)
}
