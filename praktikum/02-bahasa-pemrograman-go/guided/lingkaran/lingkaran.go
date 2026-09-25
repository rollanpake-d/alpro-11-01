package main

import "fmt"

func main() {
	var pi = 3.14
	var jarijarilingkaran, luaslingkaran float64

	//Membaca Input
	fmt.Scan(&jarijarilingkaran)

	//Menghitung Luas Lingkaran
	luaslingkaran = pi * jarijarilingkaran * jarijarilingkaran

	//Menampilkan Output
	fmt.Println(luaslingkaran)
}
