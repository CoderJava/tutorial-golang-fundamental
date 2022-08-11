package main

import "fmt"

func main() {
	luas, keliling := calculate(5, 4)
	fmt.Println("luas:", luas, "keliling:", keliling)

	luas2, keliling2 := calculate2(3, 5)
	fmt.Println("luas:", luas2, "keliling:", keliling2)
}

func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

func calculate2(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return luas, keliling
}