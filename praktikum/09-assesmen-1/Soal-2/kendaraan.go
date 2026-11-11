package main

import "fmt"

func main() {
	var kendaraan string
	jumlahMotor := 0
	jumlahMobil := 0
	jumlahTrukBus := 0

	for {
		if _, err := fmt.Scan(&kendaraan); err != nil {
			break
		}

		if kendaraan == "M" {
			jumlahMotor++
		} else if kendaraan == "K" {
			jumlahMobil++
		} else if kendaraan == "T" {
			jumlahTrukBus++
		} else {
			
			break
		}
	}

	fmt.Printf("Tipe M = %d\n", jumlahMotor)
	fmt.Printf("Tipe K = %d\n", jumlahMobil)
	fmt.Printf("Tipe T = %d\n", jumlahTrukBus)
}