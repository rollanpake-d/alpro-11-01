package main

import "fmt"

func main () {
	var kelas, hari string
	var usia int
	var harga float64

	fmt.Scan(&kelas, &usia, &hari)

	switch kelas {
	case "ekonomi" :
		harga = 100000
	case "bisnis" :
		harga = 200000
	case "eksekutif" :
		harga = 350000
	default:
		fmt.Println("Kelas tidak valid")
		return
	}

	switch hari {
	case "sabtu" :
		harga = harga + (harga * 0.2)
	case "minggu" :
		harga = harga + (harga * 0.2)
	}

	switch {
	case usia < 3:
		harga = 0
	case usia >= 3 && usia <= 12:
		harga = harga - (harga * 0.5)
	case usia >= 60:
		harga = harga - (harga * 0.3)
	}
	fmt.Printf("%.0f\n", harga)
}