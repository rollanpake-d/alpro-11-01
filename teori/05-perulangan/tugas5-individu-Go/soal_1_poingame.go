package main

import "fmt"

func main () {
	var masukkan, total int

	fmt.Scan(&masukkan)

	for masukkan != 0 {
		total += masukkan
		fmt.Scan(&masukkan)
	}
	fmt.Println(total)
}