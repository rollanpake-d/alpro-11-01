package main

import "fmt"

func main() {
	var p int
	fmt.Scan(&p)

	//blok ini hasil AI
	dewasa := p / 5
	if p%5 != 0 {
		dewasa++ 
	}
	if dewasa > 3 {
		dewasa = 3
	}

	sisaOrang := p - (dewasa * 5)
	if sisaOrang < 0 {
		sisaOrang = 0
	}

	kecil := 0
	takBerangkat := 0
	if sisaOrang > 0 {
		kecil = sisaOrang / 2
		if sisaOrang%2 != 0 {
			kecil++
		}
		
		if kecil > 5 {
			sisaOrang = sisaOrang - (5 * 2)
			kecil = 5
			takBerangkat = sisaOrang
		}
	}

	if dewasa > 0 && kecil == 0 && takBerangkat == 0 {
		fmt.Printf("Dewasa %d\n", dewasa)
	} else if dewasa > 0 && kecil > 0 && takBerangkat == 0 {
		fmt.Printf("Dewasa %d, kecil %d\n", dewasa, kecil)
	} else if takBerangkat > 0 {
		fmt.Printf("Dewasa %d, kecil %d, dan %d tak berangkat\n", dewasa, kecil, takBerangkat)
	}
}