package main

import "fmt"

func main () {
	var n int

	fmt.Scan(&n)

	var angka, max int

	fmt.Scan(&angka)

	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}