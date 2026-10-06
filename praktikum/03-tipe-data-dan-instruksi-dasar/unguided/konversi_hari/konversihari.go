package main

import "fmt"

func main() {

	var jumlahhari int

	fmt.Scan(&jumlahhari)

	tahun := jumlahhari / 360
	sisa := jumlahhari % 360

	bulan := sisa / 30
	sisa = sisa % 30

	minggu := sisa / 7
	sisa = sisa % 7

	sisahari := sisa % 7

	fmt.Println(jumlahhari, "=", tahun, "tahun", bulan, "bulan", minggu, "minggu", sisahari, "hari")
}
