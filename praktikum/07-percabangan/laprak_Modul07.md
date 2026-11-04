# <h1 align="center">Laporan Praktikum Modul [07] - [percabangan]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [if]
di if yaitu Pernyataan percabangan dasar untuk menjalankan blok kode hanya jika suatu kondisi bernilai benar (true). Jika salah (false), blok kode akan dilewati.

### B. [if-else]
if-else: Pengembangan dari if yang menyediakan dua jalur pilihan, yaitu satu blok kode dijalankan jika kondisi benar dan blok kode alternatif (else) dijalankan jika kondisi tersebut salah.

### C. [else if]
else if: Penambahan kondisi lanjutan di antara if dan else yang memungkinkan program untuk memeriksa banyak kondisi secara berurutan satu demi satu.

## Guided

### 1. [kabisat.go]

```go
package main

import "fmt"

func main () {
	var tahun int

	fmt.Scan(&tahun)

	if (tahun % 4 == 0 && tahun%100 != 0) || (tahun % 400 == 0) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}

```
#### Deskripsi
Program Go ini untuk menyelesaikan penentuan sebuah tahun itu kabisat atau bukan, dengan menggunakan if akan dicek bila tahun habis dibagi 4 hasilnya 0 dan tahun dibagi 100 hasilnya tidak sama dengan 0 atau tahun habis dibagi 400 hasilnya 0 maka itu merupakan tahun kabisat, lalu pakai else untuk mencari output tidak dengan kondisi selain yang sudah disebutkan pada if diatas.

### 2. [segitiga.go] 

```go
package main

import "fmt"

func main () {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a <= 0 || b <= 0 && c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan Segitiga")
	} else if a == b && b == c {
		fmt.Println("Segitiga sama sisi")
	} else if a == b || b == c {
		fmt.Println("Segitiga sama kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}

```
#### Deskripsi
Program Go di atas digunakan untuk menghitung sebuah nilai input sebanyak 3 variabel yang menyatakan masuk dalam kategori bukan segitiga, segitiga sama sisi, segitiga sama kaki, segitiga sembarang.

### 3. [kalkulator.go]

```go
package main 

import "fmt"

func main () {
	var a, b int
	var operator string

	fmt.Scan(&a, &operator,&b)

	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
		} else {
			fmt.Printf("%.2f\n", float64(a)/float64(b))
		}
	case "%" :
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}

```
#### Deskripsi
Program di atas adalah kalkulator sederhana berbasis teks di Go yang menerima input berupa dua angka dan satu operator matematika (+, -, *, /, %), lalu menampilkan hasil perhitungannya menggunakan struktur switch serta menangani validasi pembagian dengan nol.

### 4. [ongkir.go] 

```go
package main 

import (
	"fmt"
	"math"
)

func main () {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Melebihi batas maksimum (30 Kg)")
		return
	} 
	var tarifPerKG int

	switch zona {
	case 1:
		tarifPerKG = 10000
	case 2:
		tarifPerKG = 15000
	case 3:
		tarifPerKG = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}
	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerKG

	if layanan == "Ekspress" {
		total = total + total / 2
	}
	fmt.Println(beratBulat)
	fmt.Println(total)
}

```
#### Deskripsi
Program ini dibuat untuk perhitungan otomatisasi sesuai kriteria yang sudah ditetapkan, dimana setiap output akan menghasilkan bilangan yang telah dibulatkan keatas dan biaya tambahan sesuai kriteria biaya zona tergantung dari kategori berat, zona, dan layanan.


## Unguided

### 1. [besok.go]

```go
package main

import "fmt"

func main() {
	var t, b, th int
	fmt.Scan(&t, &b, &th)

	kabisat := (th%4 == 0 && th%100 != 0) || (th%400 == 0)

	hari := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if kabisat {
		hari[2] = 29
	}

	if b < 1 || b > 12 || t < 1 || t > hari[b] {
		fmt.Println("Tidak valid")
		return
	}

	//hasil generate AI
	fmt.Println("Valid")
	
	t, b, th = t+1, b, th
	if t > hari[b] {
		t, b = 1, b+1
		if b > 12 {
			b, th = 1, th+1
		}
	}
	fmt.Printf("%d/%d/%d\n", t, b, th)
}
```
##### output

https://github.com/rollanpake-d/alpro-11-01/blob/3855218b3f6215245cefa53e913ebfc40911921a/praktikum/07-percabangan/unguided/besok/image.png

#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan masalah perhitungan tahun kabisat yang mempunyai nilai valid atau tidak valid, program akan mengeluarkan output valid atau tidaknya diikuti baris bawahnya sesuai tanggal, bulan, dan tahun. jika syarat dari input melebihi aturan yang sudah dicantumkan pada kode maka output akan mengeluarkan "tidak valid"

### 2. [soal_c.go]

```go
package main

import "fmt"

func main () {
	var kelas, hari string
	var usia int
	var harga float64

	fmt.Scan(&kelas, &usia, &hari)

	switch kelas {
	case "ekonomi" :
		harga = 100000
	case "bisnis" :
		harga = 200000
	case "eksekutif" :
		harga = 350000
	default:
		fmt.Println("Kelas tidak valid")
		return
	}

	switch hari {
	case "sabtu" :
		harga = harga + (harga * 0.2)
	case "minggu" :
		harga = harga + (harga * 0.2)
	}

	switch {
	case usia < 3:
		harga = 0
	case usia >= 3 && usia <= 12:
		harga = harga - (harga * 0.5)
	case usia >= 60:
		harga = harga - (harga * 0.3)
	}
	fmt.Printf("%.0f\n", harga)
}

```
##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/4ac2404d6bc26a4c9014ce5aee5ef404bba373bc/praktikum/07-percabangan/unguided/soal_c/image.png

#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan masalah perhitungan tiket kereta api berdasarkan kategori kelas kereta, usia dan hari, menggunakan switch case di setiap bagian pada kelas, kereta, dan usia. dengan hasil output berupa harga dari hasil perhitungan diskon yang diperoleh, umur pembeli dan jenis kelas keretanya. 

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "besok.go"
### 2. "soal_c.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan percabangan if, if-else, dan else if. ketiganya bebas digunakan tergantung kebutuhan dan kondisi.

## Referensi
1. Modul 07_Percabangan - sumber: Modul Informatics lab,Telkom University.
2. Gemini AI