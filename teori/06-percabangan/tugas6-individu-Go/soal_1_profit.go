package main

import "fmt"

func main () {
	var bulan1, bulan2, selisih float64
	var status string

	fmt.Scan(&bulan1, &bulan2)

	if bulan1 < bulan2 {
		selisih = bulan2 - bulan1
		status = "peningkatan sebesar"
	} else if bulan1 > bulan2 {
		selisih = bulan1 - bulan2
		status = "penurunan sebesar"
	} else if bulan1 == bulan2 {
		status = "tetap"
	} 
	fmt.Println(status, selisih)
}