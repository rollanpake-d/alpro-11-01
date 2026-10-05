package main

import "fmt"

func main() {

	var x float64
	var y float64
	var f float64

	//Input
	fmt.Scan(&x)
	fmt.Scan(&y)

	//Rumus
	f = (5 * x * x) - 2*x*y + y*y*y/(x+1)

	//output
	fmt.Println(f)
}
