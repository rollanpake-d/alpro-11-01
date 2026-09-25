package main

import "fmt"

func main() {
	var a int
	var b int
	var tambah, kurang, kali, bagi, modulo int

	//Membaca Input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menghitung variabel
	tambah = a + b
	kurang = a - b
	kali = a * b
	bagi = a / b
	modulo = a % b

	//Menampilkan Output
	fmt.Println(tambah, kurang, kali, bagi, modulo)
}
