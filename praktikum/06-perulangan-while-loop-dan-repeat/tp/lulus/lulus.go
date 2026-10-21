package main

import "fmt"

func main () {
	var nilai int
	var status string

	for nilai != -1 {
		fmt.Print("Masukkan nilai anda : ")
		fmt.Scan(&nilai)
	 
	if nilai == -1 {
		status = "Program selesai"
	} else if nilai < 60 {
		status = "Anda tidak lulus"
	} else if nilai >= 60 {
		status = "Anda lulus ujian"
	}
	fmt.Println(status)
}
}