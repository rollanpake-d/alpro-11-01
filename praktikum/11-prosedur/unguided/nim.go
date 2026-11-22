package main

import "fmt"

const NMAX = 2002

type tabNIM [NMAX]int
//func cariNIM dibawah ini dibantu dari hasil generate AI
func cariNIM(T tabNIM, n, x int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == x {
			*idx = i
			return
		}
	}
}

func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabNIM
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	var x, idx int
	fmt.Scan(&x)
	cariNIM(T, n, x, &idx)

	if idx != -1 {
		var baru int
		fmt.Scan(&baru)
		ubahNIM(&T, idx, baru)

		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", x, idx+1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}