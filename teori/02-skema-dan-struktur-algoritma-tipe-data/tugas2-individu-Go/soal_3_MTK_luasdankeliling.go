package main

import "fmt"

func main() {

	var p, l int
	var luas, keliling int

	//Input
	fmt.Scan(&p)
	fmt.Scan(&l)

	//Rumus
	luas = p * l
	keliling = 2 * (p + l)

	//Output
	fmt.Println(luas, keliling)
}
