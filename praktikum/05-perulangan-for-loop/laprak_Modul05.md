# <h1 align="center">Laporan Praktikum Modul [05] - [Perulangan-for-loop]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [Perulangan]
Perulangan adalah struktur control dalam bahasa pemrograman Go, perulangan akan selalu diproses berulang disaat suatu kondisi masih terpenuhi pada satu blok kode.

#### 1. [For Loop]
Perulangan For Loop, ada 4 bentuk yaitu For dengan 3 komponen (inisialisasi, kondisi, dan post statement), lalu For hanya dengan kondisi, lalu For tanpa kondisi, serta For dengan range.

- Pada For dengan 3 komponen, ditandai dengan adanya awal inisialisasi atau variabel paling awal (belum pernah dideskripsikan) setelahnya ada kondisi biasanya ditandai dengan >=, <=, terakhir ada post statement untuk memberikan iterasi sesuai dengan ketentuan dari kondisi
- Pada For hanya dengan kondisi, penggunaan inisialisasinya digunakan untuk semua bagian pada code keseluruhan, berbeda dengan For 3 Komponen dimana setiap blok kode bisa mempunyai inisialisasi lebih dari satu.
- Pada For tanpa kondisi, menggunakan "break" untuk menghentikan program pada kondisi tertentu
- Pada For range, digunakan untuk tempat untuk mengiterasi variabel (sekumpulan data)  yang sudah ada yang nantinya akan di keluarkan secara berurutan mulai dari variabel yang tertera pada code program pertama hingga terakhir secara urut tanpa terlewat satupun

#### 2. [Nested For Loop]
Pada Nested For Loop atau For Bersarang, akan digunakan for tetapi didalam blok kode tersebut terdapat for lain, bisa terdapat 2 atau 3 bahkan lebih for didalamnya.



## Guided

### 1. [bioskop.go]

```go
package main

import "fmt"

func main () {
	var t int
	fmt.Print("input t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalpendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisfilm, hari string
		var jumlahtiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisfilm, &hari, &jumlahtiket)

		hargasatuan := 0

		if jenisfilm == "Reguler" {
			if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
				hargasatuan = 40000
			} else {
				hargasatuan = 50000
			}

		} else if jenisfilm == "3D" {
			if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
				hargasatuan = 60000
			} else {
				hargasatuan = 70000
			}

		} else if jenisfilm == "IMAX" {
			if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
				hargasatuan = 80000
			} else {
				hargasatuan = 90000
			}
		}
			totalharga := hargasatuan * jumlahtiket

			if jumlahtiket >= 5 {
				diskon := (totalharga * 10) / 100
				totalharga = totalharga - diskon
			}

			fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalharga)

			totalpendapatan = totalpendapatan + totalharga
		}

		fmt.Println("output total pendapatan:", totalpendapatan)
	}

```
#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan masalah dari sebuah bioskop yang ingin mengetahui perhitungan total pendapatan dari transaksi yang telah diinputkan transaksi melakukan perhitungan berdasarkan jenis film, hari pemutaran, dan jumlah tiket yang dibeli dimana jika setiap satu kali transaksi lebih dari 5 tiket akan diberikan diskon sebesar 10% dari total harga sebelum diskon.

### 2. [tabungan.go]

```go
package main

import "fmt"

func main () {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	total := 0;

	TabunganHariIni := a;

	for i := 1; i <= n; i++ {
		total = total + TabunganHariIni

		TabunganHariIni = TabunganHariIni + b

	}
		fmt.Println(total)

}

```
#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan masalah dari persoalan tabungan yang diinputkan seseorang dimana program akan menghitung nilai uang pertama hingga hari berikutnya yang telah ditentukan.

### 3. [terbesar.go]

```go
package main

import "fmt"

func main () {
	var n int

	fmt.Scan(&n)

	var angka, max int

	fmt.Scan(&angka)

	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}
```
#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan masalah dari persoalan pemilihan satu angka yang bernilai paling besar diantara angka yang lain, jumlah banyaknya angka akan disesuaikan menurut input yang diberikan oleh usernya.


## Unguided

### 1. [prima.go]

```go
package main

import "fmt"

func main () {
	var n int
	var status string
	fmt.Scan(&n)

	jumlahFaktor := 0

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		} 
	} 

	if jumlahFaktor == 2{
		status = "Ya"
	} else if jumlahFaktor != 2 {
		status = "Tidak"
	}
	fmt.Println(status)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/818e46248198c3ab801e582561b7788f2dc355c9/praktikum/05-perulangan-for-loop/unguided/prima/image.png

#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan persoalan sebuah pemfaktoran, dimana terdapat variabel n yang diberikan oleh user dan bagian for untuk mengecek jika variabel n habis dibagi i maka akan ditambahkan nilai 1, pada bagian bawahnya di cek bila terdapat hanya 2 bilangan maka status akan Ya jika tidak hanya 2 maka status akan Tidak

### 2. [ecct.go]

```go
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
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/3989afa3505f5c1b7b7943c2ef5a428536af4922/praktikum/05-perulangan-for-loop/unguided/ecct/image.png

#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan persoalan menghitung jumlah rata-rata mahasiswa jika nilai total yang diinputkan berada dibawah rata-rata 3.33 maka status Tidak, jika nilai total yang diinputkan berada tepat di rata-rata 3.33 atau lebih dari 3.33 maka status Ya, disini juga terdapat hasil dari jumlah status antara Ya dan Tidak, jika ada yang mendapatkan Ya maka jumlah yang mendapat sertifikasi akan menghasilkan output sesuai jumlah Ya, begitupun juga sebaliknya dengan status Tidak.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "prima.go"
### 2. "ecct.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan perulangan For dengan 4 bentuk yaitu for dengan 3 komponen(inisialisasi, kondisi, dan post statement), For hanya dengan kondisi, For tanpa kondisi, dan for dengan range serta Nested For Loop atau For bersarang

## Referensi
1. Modul 05 - Perulangan For - Loop.PDF  - sumber: Modul Informatics lab,Telkom University.