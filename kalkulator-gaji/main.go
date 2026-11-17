package main

import "fmt"

type gaji struct {
	gaji_pokok    int
	gaji_lembur   int
	potongan_gaji int
}

func inputGaji() gaji {
	var g gaji
	fmt.Print("Masukkan gaji pokok: Rp")
	fmt.Scan(&g.gaji_pokok)
	fmt.Print("Masukkan gaji lembur: Rp")
	fmt.Scan(&g.gaji_lembur)
	fmt.Print("Masukkan potongan gaji: Rp")
	fmt.Scan(&g.potongan_gaji)
	return g
}

func hitungGaji(t gaji) int {
	return t.gaji_pokok + t.gaji_lembur - t.potongan_gaji
}

func main() {
	p := inputGaji()
	fmt.Println("Total gaji: Rp", hitungGaji(p))

}
