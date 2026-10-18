package main

import "fmt"

func main () {
	var n int
	var status string
	fmt.Scan(&n)

	jumlahFaktor := 0

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		} 
	} 

	if jumlahFaktor == 2{
		status = "Ya"
	} else if jumlahFaktor != 2 {
		status = "Tidak"
	}
	fmt.Println(status)
}