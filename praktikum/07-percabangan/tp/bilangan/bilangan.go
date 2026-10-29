package main

import "fmt"

func main () {
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan % 2 == 0 && bilangan > 0 {
		fmt.Println("Genap, Positif")
	} else if bilangan % 2 == 0 && bilangan < 0 {
		fmt.Println("Genap, Negatif")
	} else if bilangan % 2 != 0 && bilangan > 0 {
		fmt.Println("Ganjil, Positif")
	} else if bilangan % 2 != 0 && bilangan < 0 {
		fmt.Println("Ganjil, Negatif")
	} else {
		fmt.Println("Nol")
	}
}