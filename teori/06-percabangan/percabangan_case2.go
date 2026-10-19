package main

import "fmt"

func main() {
	var nilai int
	var grade string

	fmt.Scan(&nilai)

	if nilai < 50 {
		grade = "E"
	} else if nilai >= 50 && nilai < 70 {
		grade = "C"
	} else if nilai >= 70 && nilai < 85 {
		grade = "B"
	} else if nilai >= 85 && nilai <= 100 {
	    grade = "A"
	}
	fmt.Printf("%s adalah Grade yang didapat berdasarkan nilai tersebut", grade)
}