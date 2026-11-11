package main

import "fmt"

//hitung adalah nama fungsi
func hitung(x, y int) int { //paramater disini berupa x dan y
	return 2*x + y - 5
}

//hasil dan proses ini adalah instruksi atau perintah atau proses
func main() {
	var hasil int
	hasil = hitung(10, 2) //argumen disini berupa 10 dan 2 
	fmt.Println("Hasil dari fungsi hitung adalah:", hasil)

	hasil = hitung(20, -3) //argumen disini berupa 20 dan -3
	fmt.Println("Hasil dari fungsi hitung adalah:", hasil)

	hasil = hitung(15, 4) //argumen disini berupa 15 dan 4
	fmt.Println("Hasil dari fungsi hitung adalah:", hasil)
}
