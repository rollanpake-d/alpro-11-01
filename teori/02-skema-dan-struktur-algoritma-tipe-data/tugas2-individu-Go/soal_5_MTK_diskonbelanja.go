package main

import "fmt"

func main() {

	var harga_awal int
	var harga_setelah_diskon int
	var potongan int
	var besar_diskon int

	//Input
	fmt.Scan(&harga_awal)
	fmt.Scan(&besar_diskon)

	//Rumus
	potongan = harga_awal * besar_diskon / 100
	harga_setelah_diskon = harga_awal - potongan

	fmt.Println(harga_setelah_diskon)
}
