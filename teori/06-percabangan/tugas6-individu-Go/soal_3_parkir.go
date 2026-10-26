package main

import "fmt"

func main () {
	var h1, m1, h2, m2 int

	fmt.Scan(&h1, &m1, &h2, &m2)

		if h2 < h1 {
			h2 = h2 + 12
		}
			totalMenitMulai := (h1 * 60) + m1
			totalMenitSelesai := (h2 * 60) + m2

			selisihMenit := totalMenitSelesai - totalMenitMulai

			hh := selisihMenit / 60
			mm := selisihMenit % 60

			fmt.Println(hh, "jam", mm, "menit")
}