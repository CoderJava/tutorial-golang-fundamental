package main

import "fmt"

func main() {
	scores := [8]int{
		100,
		80,
		75,
		92,
		70,
		93,
		88,
		67,
	}

	// hitung rata-rata
	// masukkan nilai >= 90 kedalam slice goodScores
	average := 0.0
	var goodScores []int
	for _, itemScore := range scores {
		average += float64(itemScore)
		if itemScore >= 90 {
			goodScores = append(goodScores, itemScore)
		}
	}
	average = average / float64(len(scores))
	fmt.Println("average:", average)
	fmt.Println("good scores:", goodScores)

}
