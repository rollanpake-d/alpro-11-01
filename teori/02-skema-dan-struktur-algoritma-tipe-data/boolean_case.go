package main

import "fmt"

func main() {
	var true bool
	var false bool

	fmt.Scan(&true)
	fmt.Scan(&false)

	//Rumus
	true = false
	false = true

	fmt.Println(true, false)
}
