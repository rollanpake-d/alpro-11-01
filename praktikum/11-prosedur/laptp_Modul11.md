# <h1 align="center">Tugas Pendahuluan Modul [11] - [11-prosedur]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. sekolah.go
```go
package main

import "fmt"

const NMAX = 1001

type tabMurid [NMAX]string
// func cariMurid dibawah ini dari generate AI
func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	*posisi = 0
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1
			return
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			*jumlah++
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabMurid
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	var cari string
	fmt.Scan(&cari)

	var posisi, jumlah int
	var ketemu bool

	cariMurid(T, n, cari, &posisi, &ketemu)

	var hurufAwal byte
	if len(cari) > 0 {
		hurufAwal = cari[0]
	}
	hitungHuruf(T, n, hurufAwal, &jumlah)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", hurufAwal, jumlah)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/bc6219c9686e34e73c0c3173bd18f4284484b268/praktikum/11-prosedur/tp/sekolah/image.png

#### Deskripsi
Kode Go tersebut berfungsi untuk mencari nama murid dalam daftar, mengetahui nomor urut absensinya jika terdaftar, serta menghitung berapa banyak murid yang namanya diawali huruf yang sama dengan huruf pertama nama yang dicari.

### 2. penduduk.go

```go
package main

import "fmt"

const NMAX = 1000

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

type tabPenduduk [NMAX]penduduk

func cariTahun(T tabPenduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahun == th {
			*idx = i
			return
		}
	}
}
// func tertuaTermuda dibawah ini dibantu dari generate AI
func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]
	for i := 1; i < n; i++ {
		if T[i].tahun < tua.tahun {
			*tua = T[i]
		}
		if T[i].tahun > muda.tahun {
			*muda = T[i]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var T tabPenduduk
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}

	var th int
	fmt.Scan(&th)

	var idx int
	cariTahun(T, n, th, &idx)

	var tua, muda penduduk
	if n > 0 {
		tertuaTermuda(T, n, &tua, &muda)
	}

	if idx != -1 {
		fmt.Printf("%s %d %s ditemukan di index ke-%d\n", T[idx].nama, T[idx].tahun, T[idx].kota, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/bc6219c9686e34e73c0c3173bd18f4284484b268/praktikum/11-prosedur/tp/penduduk/image.png

#### Deskripsi
kode ini berfungsi untuk menyimpan data penduduk (nama, tahun lahir, dan kota), mencari penduduk pertama berdasarkan tahun lahir tertentu, serta menentukan penduduk tertua dan termuda dari seluruh data yang dimasukkan.

## Kesimpulan
kedua program tersebut berfungsi untuk mengelola dan mengolah daftar data (nama murid serta data penduduk) guna mencari informasi tertentu, menghitung jumlah berdasarkan huruf awal, dan mengetahui data paling tua atau paling muda.