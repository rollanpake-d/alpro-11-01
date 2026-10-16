# <h1 align="center">Laporan Praktikum Modul [05] - [Perulangan-for-loop]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

#### A. [Pengertian Runtunan / Sekuensi]
Runtunan atau sekuensi merupakan struktur kontrol dalam algoritma pemrograman yang melakukan proses bertahap dari perintah bagian atas hingga bagian bawah secara berurutan tanpa melewatkan satu perintah, sehingga apabila ada perpindahan perintah maka output bisa terganti atau justru menghasilkan error.

#### B. [Ciri-Ciri Runtunan / Sekuensi]
1. Program diproses secara urut dari paling atas hingga paling bawah tanpa melewatkan satu perintah pun.
2. Ketika ada perpindahan/pertukaran perintah maka bisa menyebabkan output yang berbeda bahkan bisa error.
3. Tidak bisa menggunakan perulangan atau percabangan.
4. Perintah hanya bisa dijalankan satu kali saja.

#### C. [Kegunaan Runtunan / Sekuensi]
1. Untuk menampilkan Biodata.
2. Untuk mengkonversi rumus matematis yang sederhana (luas, keliling, konversi satuan)
3. Menukar 2 nilai dua variabel



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
Pada tugas guided ini diberikan soal yang berfungsi untuk memberikan hasil nama serta grade berdasarkan input yang diberikan dari user, menggunakan else if serta package bufio dan os yang baru pernah digunakan pada tugas kali ini.

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
Pada tugas guided kali ini diberikan soal untuk memberikan output keterangan berdasarkan klasifikasi dari input usia dan gaji dengan menggunakan switch case, dimana hal ini bisa dikatakan lebih simple dari model else if karena switch case tidak perlu penambahan tanda kurung "()"


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
Pada soal guided kali ini diberikan persoalan untuk menyelesaikan masalah menentukan grade siswa berdasarkan input nilai, dimana sistem ini akan menyeleksi bagian yang sesuai dengan ketentuan yang ada.


## Unguided

### 1. [grade.go]

```go
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "F"
	}
	
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/modif_grade/Screenshot%202026-10-11%20133008.png

#### Deskripsi
Pada soal unguided kali ini diberikan persoalan untuk memodifikasi dari soal guided 1 (grade.go) untuk menggunakan model switch case (bukan else if) bisa dibilang switch case ini lebih simple karena tidak perlu ada penambahan tanda kurung "()"

### 2. [pajak.go]

```go
package main

import "fmt"

func main() {
	var pajak, penghasilan float64

	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
	 pajak = penghasilan * 0.05 

	} else if penghasilan > 50 && penghasilan <= 100 {
		pajak = (0.05 * 50) + (0.1 * (penghasilan - 50)) 

	} else if penghasilan > 100 && penghasilan <= 200 {
		pajak = (0.05 * 50) + (0.1 * 50) + (0.15 * (penghasilan - 100))
	
	} else if penghasilan > 200 {
		pajak = (0.05 * 50) + (0.1 * 50) + (0.15 * 100) + (0.2 * (penghasilan - 200))
	}
	fmt.Println(pajak)
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/pajak/Screenshot%202026-10-11%20181946.png

#### Deskripsi
Pada soal unguided kali ini diberikan persoalan untuk membuat sebuah program untuk menghitung pajak berdasarkan gaji, dimana setiap kelipatan 50 juta akan ada kenaikan 5% dari nilai pajak awal, berfungsi untuk menentukan besaran output yang lebih akurat sesuai dengan rumus ketentuan yang ada.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "modif_grade.go"
### 2. "pajak.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan else if, switch case serta packages bufio, os, dan algoritma yang lebih komples namun harus lebih terstruktur.

## Referensi
1. Modul 04 - Runtunan / Sekuensi.PDF  - sumber: Modul Informatics lab,Telkom University.
