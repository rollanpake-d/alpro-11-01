package main

import "fmt"

func main() {
	var totalBelanja int
	var kartu bool

	fmt.Scan(&totalBelanja, &kartu)

	statusKartu := kartu
	diskon := false
	cashback := false

	if totalBelanja >= 100000 {
		diskon = true
	}

	if totalBelanja >= 200000 && kartu {
		cashback = true
	}

	hargaAkhir := float64(totalBelanja)

	if diskon {
		hargaAkhir = hargaAkhir - (hargaAkhir * 0.10)
	}

	if cashback {
		hargaAkhir = hargaAkhir - 75000
	}

	fmt.Println("Kartu?", statusKartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Printf("Rp. %.0f\n", hargaAkhir)
}