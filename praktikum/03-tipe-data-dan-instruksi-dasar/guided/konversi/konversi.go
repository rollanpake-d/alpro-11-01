package main

import "fmt"

func main() {

	var celcius float64

	fmt.Print("Masukkan Suhu: ")
	fmt.Scanln(&celcius)

	fmt.Println(celcius + 273)
}
