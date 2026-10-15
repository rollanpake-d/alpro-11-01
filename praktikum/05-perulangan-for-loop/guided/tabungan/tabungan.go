package main

import "fmt"

func main () {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	total := 0;

	TabunganHariIni := a;

	for i := 1; i <= n; i++ {
		total = total + TabunganHariIni

		TabunganHariIni = TabunganHariIni + b

	}
		fmt.Println(total)

}
