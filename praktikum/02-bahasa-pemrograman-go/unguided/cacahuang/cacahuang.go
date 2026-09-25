package main

import "fmt"

func main() {
	var NominalUang int

	//Membaca Input
	fmt.Scan(&NominalUang)

	//Menghitung Cacah Uang
	Lembar10k := NominalUang / 10000
	sisa := NominalUang % 10000

	Lembar5k := sisa / 5000
	sisa = sisa % 5000

	Lembar1k := sisa / 1000

	//Menampilkan Output
	fmt.Println(Lembar10k, Lembar5k, Lembar1k)
}
