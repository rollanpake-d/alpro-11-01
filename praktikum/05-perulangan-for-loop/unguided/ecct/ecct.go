package main

import "fmt"

func main () {
	var m int
	var P, V, GA, FS, CI, TC int
	var status string

	fmt.Scan(&m)

	jumlahSertifikat := 0

	for i := 1; i <= m; i++ {

		fmt.Scan(&P, &V, &GA, &FS, &CI, &TC)
		rataRata := float64(P + V + GA + FS + CI + TC) / 6.0

	if rataRata < 3.33 {
	   status = "Tidak"
	} else if rataRata >= 3.33 {
	   status = "Ya"
	   jumlahSertifikat++
	}
	fmt.Println(status)
}
	fmt.Println(jumlahSertifikat)
}