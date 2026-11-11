package main

import "fmt"

func main() {
	var jumlahDaun int
	if _, err := fmt.Scan(&jumlahDaun); err != nil || jumlahDaun <= 0 {
		return
	}

	var lebarTerkecil int
	for i := 0; i < jumlahDaun; i++ {
		var lebar int
		fmt.Scan(&lebar)

		if i == 0 || lebar < lebarTerkecil {
			lebarTerkecil = lebar
		}
	}
	fmt.Println(lebarTerkecil)
}