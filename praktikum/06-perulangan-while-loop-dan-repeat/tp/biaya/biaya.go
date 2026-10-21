package main

import "fmt"

func main () {
	var n int
	var biayaKuliah float64
	
fmt.Scan(&n)

biayaKuliah = 10000000

	for i := 2; i <= n; i++ {
		biayaKuliah = biayaKuliah * 1.06
	}
	fmt.Printf("%.0f", biayaKuliah)
}