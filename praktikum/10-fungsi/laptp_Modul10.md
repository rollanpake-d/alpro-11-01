# <h1 align="center">Tugas Pendahuluan Modul [10] - [10-fungsi]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. Jelaskan perbedaan antara if,if-else, else if, dan switch pada bahasa pemrograman Go! Sebutkan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java. 

Dalam bahasa Go, if, if-else, dan else if digunakan untuk percabangan berbasis kondisi boolean bertingkat, sementara switch dirancang untuk pencocokan nilai atau kondisi yang lebih rapi tanpa memerlukan break otomatis seperti pada Java.

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
Pada soal kali ini digunakan untuk menentukan status bilangan genap atau ganjil dan positif atau negatif, menggunakan else if untuk menentukan ganjil,genap,positif, dan negatif serta else untuk menentukan nilai diluar itu yaitu 0.

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
pada kode kali ini digunakan untuk menyelesaikan masalah penentuan nama hari berdasarkan bilangan bulat yang diinput sesuai dengan urutan ke satu yaitu senin hingga urutan ketujuh yaitu minggu menggunakan switch case untuk membungkus input berupa angka untuk menyatakan variabel nomorHari dan memakai default untuk mengeluarkan "Nomor hari tidak valid" bila user memberikan input angka lebih dari 7.

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan if, else if, if-else, dan switch case