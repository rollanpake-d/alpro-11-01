package main

import "fmt"

func luasSegitiga (alas, tinggi int, luas*float64) {
	*luas = float64(alas*tinggi) / 2
}

func main () {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var alas, tinggi int
		fmt.Scan(&alas, &tinggi)
		var luas float64
		luasSegitiga(alas, tinggi, &luas)
		fmt.Println(luas)
	}
}