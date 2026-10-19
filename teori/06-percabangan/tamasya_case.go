package main

import "fmt"

func main () {
	var N, mobil uint8
	fmt.Scan(&N)

	mobil = N / 7 // 10 / 7 -> 3

	if N%7 != 0 { // N ->10, apakah 10 mod 7 != 0 -> true
		mobil++	
	}
	fmt.Println(mobil)
}