package main

import "fmt"

func main() {

	var x int

	fmt.Println("Masukkan Nominal: ")
	fmt.Scan(&x)

	var SepuluhRibu int = x / 10000
	var sisa int = x % 10000

	var LimaRibu int = sisa / 5000
	sisa = sisa % 5000

	var Seribuan int = sisa / 1000

	fmt.Println(SepuluhRibu, LimaRibu, Seribuan)
}
