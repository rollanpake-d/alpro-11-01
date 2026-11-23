package main

import "fmt"

func cekSama(hasil *bool, a, b int) {
	*hasil = a == b
	fmt.Println(*hasil)
}

// a = 10 kemudian b = 5 maka di false
// a = 10 kemudian b = 10 maka di true

func main() {
	var a, b int
	var hasil bool
	a = 10
	b = 10
	cekSama(&hasil, a, b) //prosedur dengan pass by references untuk hasil dan pass by value untuk a dan b
	hasil = false
	fmt.Println(hasil)
}