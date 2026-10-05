package main

import "fmt"

func main() {

	var x float64

	//Input
	fmt.Scan(&x)

	//Rumus
	(x) = (x*x + 2*x + 1) / (x - 3)

	//Ouput
	fmt.Println(x)
}
