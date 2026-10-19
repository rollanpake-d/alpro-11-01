package main

import "fmt"

func main () {
	var belanja, potongan float64
	var assesmen bool

	fmt.Scan(&belanja, &assesmen)

	if assesmen {
		potongan = belanja * 0.35
		belanja = belanja - potongan
	}
	fmt.Println(belanja)
}