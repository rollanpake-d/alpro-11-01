package main 

import (
	"fmt"
	"math"
)

func main () {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Melebihi batas maksimum (30 Kg)")
		return
	} 
	var tarifPerKG int

	switch zona {
	case 1:
		tarifPerKG = 10000
	case 2:
		tarifPerKG = 15000
	case 3:
		tarifPerKG = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}
	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerKG

	if layanan == "Ekspress" {
		total = total + total / 2
	}
	fmt.Println(beratBulat)
	fmt.Println(total)
}