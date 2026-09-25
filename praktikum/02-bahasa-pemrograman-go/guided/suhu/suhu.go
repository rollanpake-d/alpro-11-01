package main

import "fmt"

func main() {
	var celcius float64
	var reamur, fahrenheit, kelvin float64

	//Membaca Input
	fmt.Scan(&celcius)

	//Menghitung Suhu
	reamur = celcius * 4.0 / 5.0
	fahrenheit = celcius*9.0/5.0 + 32.0
	kelvin = celcius + 273.15

	//Menampilkan output
	fmt.Println(reamur)
	fmt.Println(fahrenheit)
	fmt.Println(kelvin)

}
