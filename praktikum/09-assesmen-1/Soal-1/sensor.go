package main

import "fmt"

func main() {
	var tertinggi, terendah, hitung int
	if _, err := fmt.Scan(&tertinggi); err != nil {
		return
	}

	maks, min, jumlah, hitung := tertinggi, tertinggi, tertinggi, 1

	for fmt.Scan(&terendah); terendah != tertinggi; fmt.Scan(&terendah) {
		if terendah > maks { maks = terendah }
		if terendah < min { min = terendah }
		jumlah += terendah
		hitung++
		tertinggi = terendah
	}

	fmt.Printf("%d %d %.2f\n", maks, min, float64(jumlah)/float64(hitung))
}