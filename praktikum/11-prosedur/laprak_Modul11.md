# <h1 align="center">Laporan Praktikum Modul [11] - [prosedur]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [deskripsi prosedur]
Prosedur adalah subprogram atau blok kode yang dirancang untuk menjalankan serangkaian instruksi atau aksi tertentu tanpa mengembalikan nilai (return value) secara langsung kepada pemanggilnya.

## Guided

### 1. [sapi.go]

```go
package main

import "fmt"

const NMAX = 10000

type Arr [NMAX]int

func mengisiSapi(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func geserKanan(T *Arr, n int) {
	var temp int = T[n-1]

	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}

	T[0] = temp
}

func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var T Arr
	mengisiSapi(&T, n)

	k = k % n

	for i := 0; i < k; i++ {
		geserKanan(&T, n)
	}

	cetak(T, n)

}

```
#### Deskripsi
Kode Go tersebut berfungsi untuk mengeser (merotasi) elemen-elemen di dalam sebuah array ke arah kanan secara melingkar (right circular shift) sebanyak k kali, lalu mencetak hasilnya.

Ringkasan Alur

## Unguided

### 1. [nim.go]

```go
package main

import "fmt"

const NMAX = 2002

type tabNIM [NMAX]int
//func cariNIM dibawah ini dari hasil generate AI
func cariNIM(T tabNIM, n, x int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == x {
			*idx = i
			return
		}
	}
}

func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabNIM
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	var x, idx int
	fmt.Scan(&x)
	cariNIM(T, n, x, &idx)

	if idx != -1 {
		var baru int
		fmt.Scan(&baru)
		ubahNIM(&T, idx, baru)

		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", x, idx+1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}
```
##### output

https://github.com/rollanpake-d/alpro-11-01/blob/fa05509e753da711de62f86cf2473cde3105d145/praktikum/11-prosedur/unguided/Screenshot%202026-11-22%20100306.png

#### Deskripsi
Kode Go tersebut berfungsi untuk mencari suatu NIM di dalam daftar NIM mahasiswa, mengubah nilainya dengan NIM baru jika ditemukan, lalu menampilkan posisi awal serta daftar NIM terbaru.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "nim.go"

## Kesimpulan
Tujuan praktikum kali ini adalah mengenalkan konsep blok kode prosedur serta penggunaan parameter untuk menjalankan aksi atau memodifikasi data tanpa menggunakan nilai kembalian (return value).

## Referensi
1. Modul 11_Prosedur - sumber: Modul Informatics lab,Telkom University.
2. Gemini AI