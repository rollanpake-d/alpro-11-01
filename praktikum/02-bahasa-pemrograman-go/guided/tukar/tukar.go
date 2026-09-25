package main

import "fmt"

func main() {
	var a, b int

	//Membaca Input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar Nilai
	a, b = b, a

	//Menampilkan Output
	fmt.Println(a)
	fmt.Println(b)
}
