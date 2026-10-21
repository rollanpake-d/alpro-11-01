package main

import "fmt"

func main () {
	var n int

	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Print("Genap, ")
	} else {
		fmt.Print("Ganjil, ")
	}

	if n > 0 {
		fmt.Println("Positif") 
	} else if n < 0 {
		fmt.Println("Negatif")
	} else {
		fmt.Println("Nol")
	}
}