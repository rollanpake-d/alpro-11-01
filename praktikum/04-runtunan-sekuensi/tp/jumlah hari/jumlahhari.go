package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah suatu tahun merupakan tahun kabisat
func Kabisat(tahun int) bool {
	return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int
	var bulan string

	// Membaca input tahun dan nama bulan
	fmt.Scan(&tahun, &bulan)

	// Mengecek jumlah hari berdasarkan nama bulan yang valid
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if Kabisat(tahun) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		// Jika nama bulan tidak sesuai format/validasi
		fmt.Println("-")
	}
}