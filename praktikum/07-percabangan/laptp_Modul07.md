# <h1 align="center">Tugas Pendahuluan Modul [7] - [07-percabangan]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. Jelaskan perbedaan antara if,if-else, else if, dan switch pada bahasa pemrograman Go! Sebutkan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java. 

Dalam bahasa Go, if, if-else, dan else if digunakan untuk percabangan berbasis kondisi boolean bertingkat, sementara switch dirancang untuk pencocokan nilai atau kondisi yang lebih rapi tanpa memerlukan break otomatis seperti pada Java.

### 2. bilangan.go
```go
package main

import "fmt"

func main () {
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan % 2 == 0 && bilangan > 0 {
		fmt.Println("Genap, Positif")
	} else if bilangan % 2 == 0 && bilangan < 0 {
		fmt.Println("Genap, Negatif")
	} else if bilangan % 2 != 0 && bilangan > 0 {
		fmt.Println("Ganjil, Positif")
	} else if bilangan % 2 != 0 && bilangan < 0 {
		fmt.Println("Ganjil, Negatif")
	} else {
		fmt.Println("Nol")
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/ce2bbc0c8a1cc68a519bb70143f4e53b6719cd08/praktikum/07-percabangan/tp/bilangan/image.png

#### Deskripsi
Pada soal kali ini digunakan untuk menentukan status bilangan genap atau ganjil dan positif atau negatif, menggunakan else if untuk menentukan ganjil,genap,positif, dan negatif serta else untuk menentukan nilai diluar itu yaitu 0.

### 3. hari.go

```go
package main 

import "fmt"

func main () {
	var nomorHari string
	
	fmt.Scan(&nomorHari)

	switch nomorHari {
	case "1": {
		fmt.Println("Senin")
	}
	case "2": {
		fmt.Println("Selasa")
	}
	case "3": {
		fmt.Println("Rabu")
	}
	case "4": {
		fmt.Println("Kamis")
	}
	case "5": {
		fmt.Println("Jumat")
	}
	case "6": {
		fmt.Println("Sabtu")
	}
	case "7": {
		fmt.Println("Minggu")
	}
	default:
		fmt.Println("Nomor hari tidak valid ")
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/ce2bbc0c8a1cc68a519bb70143f4e53b6719cd08/praktikum/07-percabangan/tp/hari/image.png

#### Deskripsi
pada kode kali ini digunakan untuk menyelesaikan masalah penentuan nama hari berdasarkan bilangan bulat yang diinput sesuai dengan urutan ke satu yaitu senin hingga urutan ketujuh yaitu minggu menggunakan switch case untuk membungkus input berupa angka untuk menyatakan variabel nomorHari dan memakai default untuk mengeluarkan "Nomor hari tidak valid" bila user memberikan input angka lebih dari 7.

### 4. terbesar.go

```go
package main 

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println(a)
	} else if b >= a && b >= c {
		fmt.Println(b)
	} else if c >= b && c >= a {
		fmt.Println(c)
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/ce2bbc0c8a1cc68a519bb70143f4e53b6719cd08/praktikum/07-percabangan/tp/terbesar/image.png


#### Deskripsi
pada kode kali ini digunakan untuk menentukan nilai terbesar dari 3 angka yang ada, tidak menggunakan perulangan dan max, hanya pakai else if.

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan if, else if, if-else, dan switch case