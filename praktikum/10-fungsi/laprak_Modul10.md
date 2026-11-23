# <h1 align="center">Laporan Praktikum Modul [10] - [fungsi]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [deskripsi fungsi]
Fungsi adalah blok kode yang menjalankan tugas tertentu dan dapat menghasilkan sebuah nilai (return value). Nilai tersebut bisa disimpan atau digunakan kembali di bagian kode lainnya, dengan input yang dapat disesuaikan melalui parameter.

### B. [struktur penulisan fungsi] 
nama fungsi berisi daftar parameter, kamus lokal berisi deklarasi tambahan, algoritma berisi langkah perhitungan 

## Guided

### 1. [kabisat.go]

```go
package main

import "fmt"

func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}

func jumlahHari(bulan, tahun int) int {
	var hari int
	if bulan == 2 {
		if kabisat(tahun) {
			hari = 29
		} else {
			hari = 28
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		hari = 30
	} else {
		hari = 31
	}
	return hari
}

func hariKe(tanggal, bulan, tahun int) int {
	var total int = 0
	var i int
	for i = 1; i < bulan; i++ {
		total = total + jumlahHari(i, tahun)
	}
	total = total + tanggal
	return total
}

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	urutanHari := hariKe(tanggal, bulan, tahun)

	var totalHariTahun int
	if kabisat(tahun) {
		totalHariTahun = 366
	} else {
		totalHariTahun = 365
	}

	sisaHari := totalHariTahun - urutanHari

	fmt.Printf("Hari ke-%d\n", urutanHari)
	fmt.Printf("Sisa hari dalam tahun %d: %d\n", tahun, sisaHari)
}

```
#### Deskripsi
Program Go ini untuk menyelesaikan penentuan sebuah tahun itu kabisat atau bukan, dengan menggunakan func yang akan dicek bila tahun habis dibagi 4 hasilnya 0 dan tahun dibagi 100 hasilnya tidak sama dengan 0 atau tahun habis dibagi 400 hasilnya 0 maka itu merupakan tahun kabisat.

### 2. [player.go] 

```go
package main

import "fmt"

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

func idxTertinggi(T tabPlayer, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idxMax].score {
			idxMax = i
		}
	}
	return idxMax
}

func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabPlayer

	fmt.Scan(&n)
	if n < 1 {
		n = 1
	} else if n > NMAX {
		n = NMAX
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	avg := rataRata(T, n)
	bestIdx := idxTertinggi(T, n)
	aboveCount := hitungDiAtas(T, n, avg)

	fmt.Printf("Rata-rata skor: %.2f\n", avg)
	fmt.Printf("Player terbaik: %s (%d)\n", T[bestIdx].nickname, T[bestIdx].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", aboveCount)
}

```
#### Deskripsi
Program Go di atas digunakan untuk mengolah data skor sejumlah player guna menghitung rata-rata skor, menemukan player dengan skor tertinggi, serta menghitung berapa banyak player yang skornya berada di atas rata-rata tersebut.

## Unguided

### 1. [gempa.go]

```go
package main

import "fmt"

const NMAX = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m < 6.0 {
		return "Sedang"
	} else if m < 7.0 {
		return "Kuat"
	}
	return "Besar"
}

func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}
// func idxTerkuat dibawah ini dibantu dari generate AI
func idxTerkuat(T tabGempa, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idxMax].magnitudo {
			idxMax = i
		}
	}
	return idxMax
}

func hitungKategori(T tabGempa, n int, k string) int {
	jumlah := 0
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabGempa
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", T[i].nama, T[i].tahun, T[i].magnitudo, kategori(T[i].magnitudo))
	}

	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataMagnitudo(T, n))

	terkuat := T[idxTerkuat(T, n)]
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", terkuat.nama, terkuat.tahun, terkuat.magnitudo)

	fmt.Printf("Jumlah gempa kategori Besar: %d\n", hitungKategori(T, n, "Besar"))
}
```
##### output

https://github.com/rollanpake-d/alpro-11-01/blob/1026db70edb81d98d9010468d0f76c6a58f045aa/praktikum/10-fungsi/unguided/image.png

#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan pengolah data daftar gempa guna mengelompokkan tiap gempa ke dalam kategori magnitudo, menghitung rata-rata magnitudo, mencari gempa terkuat, serta menghitung jumlah gempa pada kategori tertentu (seperti kategori "Besar").

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "gempa.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan func dan return sebagai nilai balik 
## Referensi
1. Modul 10_fungsi - sumber: Modul Informatics lab,Telkom University.