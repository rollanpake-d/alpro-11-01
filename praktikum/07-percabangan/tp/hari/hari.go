package main 

import "fmt"

func main () {
	var nomorHari string
	
	fmt.Scan(&nomorHari)

	switch nomorHari {
	case "1": {
		fmt.Println("Senin")
	}
	case "2": {
		fmt.Println("Selasa")
	}
	case "3": {
		fmt.Println("Rabu")
	}
	case "4": {
		fmt.Println("Kamis")
	}
	case "5": {
		fmt.Println("Jumat")
	}
	case "6": {
		fmt.Println("Sabtu")
	}
	case "7": {
		fmt.Println("Minggu")
	}
	default:
		fmt.Println("Nomor hari tidak valid ")
	}
}