package main

import "fmt"

func main() {
	persegi := Persegi{
		Sisi: 4,
	}
	persegiPanjang := PersegiPanjang{
		Panjang: 6,
		Lebar:   5,
	}
	luasPersegi := SeberapaLuas1(persegi)
	luasPersegiPanjang := SeberapaLuas2(persegiPanjang)
	fmt.Println("luas persegi:", luasPersegi)
	fmt.Println("luas persegi panjang:", luasPersegiPanjang)
}

// Masalah jika tidak pakai interface maka, kita harus membuat satu persatu function SeberapaLuas sesuai
// dengan parameteter struct-nya masing-masing.
func SeberapaLuas1(persegi Persegi) int {
	return persegi.Sisi
}

func SeberapaLuas2(persegiPanjang PersegiPanjang) int {
	return persegiPanjang.HitungLuas()
}

// type Luas interface {
// 	HitungLuas() int
// }

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}
