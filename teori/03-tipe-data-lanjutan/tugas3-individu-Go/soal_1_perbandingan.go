package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	lebihbesar := a > b
	samadengan := a == b
	lebihkecil := a < b

	fmt.Println(lebihbesar)
	fmt.Println(samadengan)
	fmt.Println(lebihkecil)
}
