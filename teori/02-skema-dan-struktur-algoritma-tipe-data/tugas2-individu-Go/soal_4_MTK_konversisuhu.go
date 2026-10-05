package main

import "fmt"

func main() {

	var C float64
	var fahrenheit float64

	//Input
	fmt.Scan(&C)

	//Rumus
	fahrenheit = C*9/5 + 32

	//Output
	fmt.Println(fahrenheit)
}
