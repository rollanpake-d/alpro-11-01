package main

import "fmt"

func add(a, b int32) int32 { //fungsi punya keyword tipe data dan return
	return a + b
}

func eating () { // prosedur punya keyword void dan tidak return
	fmt.Println("saya lapar")
	fmt.Println("saya ambil piring")
	fmt.Println("saya ambil nasi dan beberapa lauk")
	fmt.Println("saya makan nasi dan lauk")
}

//blok func hitung ini berupa fungsi
func hitungF(x, y int) int {
	var hasil int
	hasil = 2 * x + y - 5
	return hasil
}

//blok func hitung ini berupa prosedur
func hitungFF(x, y int) {
	var hasil int
	hasil = 2 * x + y - 5
	fmt.Println(hasil)
}

func main () {
	fmt.Println(add(2, 3)) // cara panggil fungsi
	eating () // cara panggil procedur
	fmt.Println(hitungF(2, 3)) // cara panggil fungsi
	hitungFF(2, 3) // cara panggil prosedur

}