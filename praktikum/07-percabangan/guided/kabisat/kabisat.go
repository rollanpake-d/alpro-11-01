package main

import "fmt"

func main () {
	var tahun int

	fmt.Scan(&tahun)

	if (tahun % 4 == 0 && tahun%100 != 0) || (tahun % 400 == 0) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}