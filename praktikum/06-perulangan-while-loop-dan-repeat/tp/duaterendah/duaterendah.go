package main

import "fmt"

func main() {
	var n, nilai int
	var namaSiswa string
	var namaMin1, namaMin2 string

	min1 := 99999
	min2 := 99999

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&namaSiswa, &nilai)

		//bagian blok if dan else if saya ambil dari hasil AI
		if nilai < min1 {
			min2 = min1
			namaMin2 = namaMin1
			min1 = nilai
			namaMin1 = namaSiswa
		} else if nilai < min2 {
			min2 = nilai
			namaMin2 = namaSiswa
		}
	}
	fmt.Println(namaMin1)
	fmt.Println(namaMin2)
}