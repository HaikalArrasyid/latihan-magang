package main

import "fmt"

type gaji struct {
	gaji_pokok    int
	gaji_lembur   int
	potongan_gaji string
}

func inputGaji() gaji {
	var g gaji

	fmt.Print("Masukkan gaji pokok: Rp")
	fmt.Scan(&g.gaji_pokok)

	fmt.Print("Masukkan gaji lembur: Rp")
	fmt.Scan(&g.gaji_lembur)

	fmt.Print("Apakah ada potongan gaji? (y/n): ")
	fmt.Scan(&g.potongan_gaji)

	return g
}

func hitungGaji(g gaji) int {
	totalGaji := g.gaji_pokok + g.gaji_lembur

	if g.potongan_gaji == "y" {
		totalGaji -= totalGaji * 10 / 100
	} else if g.potongan_gaji == "n" {
		return totalGaji
	} else {
		fmt.Println("Input tidak valid, tidak ada potongan yang diterapkan")
	}
	return totalGaji
}

func main() {
	p := inputGaji()
	fmt.Println("Total gaji: Rp", hitungGaji(p))


}
