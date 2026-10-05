package main

import "fmt"

func main() {
	var x, low, high int

	fmt.Scan(&x, &low, &high)
	fmt.Println(x >= low && x <= high)
}
