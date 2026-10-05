package main

import "fmt"

func main() {
	var banyakbilangan, bilangan, angka_ke1, angka_ke4, hasil int

	fmt.Scan(&banyakbilangan)
	for bilangan = 1; bilangan <= banyakbilangan; bilangan++ {
		fmt.Scan(&bilangan)
		angka_ke1 = bilangan / 1000
		angka_ke4 = bilangan % 10
		hasil += angka_ke1 + angka_ke4
	}
	fmt.Println(hasil)
}
