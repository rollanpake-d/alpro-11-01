package main

import "fmt"

func main() {
	var x float64

	//Input
	fmt.Scan(&x)

	//Rumus
	x = x*x*x + 3*x/x*x*x*x - 3*x*x + 4

	//Output
	fmt.Println(x)
}
