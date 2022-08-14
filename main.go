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
	luasPersegi := SeberapaLuas(persegi)
	luasPersegiPanjang := SeberapaLuas(persegiPanjang)
	fmt.Println("luas persegi:", luasPersegi)
	fmt.Println("luas persegi panjang:", luasPersegiPanjang)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}

type BangunDatar interface {
	HitungLuas() int
}

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
