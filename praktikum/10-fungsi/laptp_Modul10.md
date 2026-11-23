# <h1 align="center">Tugas Pendahuluan Modul [10] - [10-fungsi]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. raport.go
```go
package main

import "fmt"

const NMAX = 999

type rapot struct {
	matkul string
	nilai  int
}

type tabRapot [NMAX]rapot

func nilaiHuruf(x int) string {
	if x > 80 {
		return "A"
	} else if x > 70 {
		return "AB"
	} else if x > 65 {
		return "B"
	} else if x > 60 {
		return "BC"
	} else if x > 50 {
		return "C"
	} else if x > 40 {
		return "D"
	}
	return "E"
}

func rataRata(T tabRapot, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].nilai
	}
	return float64(total) / float64(n)
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabRapot
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].matkul, &T[i].nilai)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %s\n", T[i].matkul, T[i].nilai, nilaiHuruf(T[i].nilai))
	}

	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/b0edb770e83894756344b7332b768941ad65b61a/praktikum/10-fungsi/tp/raport/image.png

#### Deskripsi
Pada soal kali ini digunakan untuk mengolah data nilai sejumlah mata kuliah guna mengonversi tiap nilai angka menjadi nilai huruf serta menghitung rata-rata seluruh nilai tersebut.

### 2. inputnilai.go

```go
package main

import "fmt"

const NMAX = 100

type student struct {
	nim   string
	name  string
	grade float64
}

type students [NMAX]student

func cariNIM(T students, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}

func rataRata(T students, n int) float64 {
	if n == 0 {
		return 0.0
	}
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}

func main() {
	var T students
	var n int
	var nim, name string
	var grade float64

	for {
		fmt.Scan(&nim)
		if nim == "STOP" {
			break
		}
		fmt.Scan(&name, &grade)

		if cariNIM(T, n, nim) == -1 && n < NMAX {
			T[n] = student{nim: nim, name: name, grade: grade}
			n++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].name, T[i].grade)
	}

	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/b0edb770e83894756344b7332b768941ad65b61a/praktikum/10-fungsi/tp/input_nilai/image.png

#### Deskripsi
Kode ini berfungsi untuk menerima dan menyimpan data mahasiswa (NIM, nama, dan nilai) secara berulang hingga input "STOP" dengan mencegah duplikasi NIM, menampilkan daftar data yang tersimpan, serta menghitung rata-rata nilainya.

## Kesimpulan
Tujuan Tugas Pendahulan ini untuk mengenalkan aturan func dan return sebagai nilai balik. 