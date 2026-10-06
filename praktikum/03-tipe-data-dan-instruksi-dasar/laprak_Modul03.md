# <h1 align="center">Laporan Praktikum Modul [03] - [tipe-data-dan-instruksi-dasar-go]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [Tipe data Integer, Real, Boolean, Char]
Pada bagian ini menjelaskan pengertian tipe data Integer, ReaL, Boolean, dan Char serta menjelaskan fungsi dan beberapa kategori dalam tipe data integer, real, dan character 

#### 1. [Judul Sub-topik 1, Integer]
Pada integer ada int8, int16, int32, dan int64 kemudian uint8, uint16, uint32, dan uint64, pada kategori int dan uint tersebut berbeda pada rentang bitnya. tipe ini untuk menyimpan bilangan bulat ditandai dengan adanya tipe data "int".

#### 2. [Judul Sub-topik 2, Real]
Pada real ada float32 dan float64. tipe ini untuk menyimpan bilangan pecahan ditandai dengan adanya tipe data "float" diikuti dengan kategori bit yang sedang dibutuhkan.

#### 3. [Judul Sub-topik 3, Boolean]
Pada boolean tipe ini untuk menyimpan variabel logika (true & false) ditandai dengan adanya tipe data "bool" 

#### 4. [Judul Sub-topik 4, Character]
Pada character ada tipe data byte, rune, dan string. tipe ini untuk menyimpan variabel sekumpulan teks atau karakter ditandai dengan adanya tipe data "char".

## Guided

### 1. [kasir.go]

```go
package main

import "fmt"

func main() {

	var x int

	fmt.Println("Masukkan Nominal: ")
	fmt.Scan(&x)

	var SepuluhRibu int = x / 10000
	var sisa int = x % 10000

	var LimaRibu int = sisa / 5000
	sisa = sisa % 5000

	var Seribuan int = sisa / 1000

	fmt.Println(SepuluhRibu, LimaRibu, Seribuan)
}
```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel x dalam tipe data integer, memberikan input variabel dengan fmt.Scan dari user, lalu menuliskan rumus pecahan uang menggunakan pembagian bilangan bulat serta modulus dan menghasilkan output sesuai dengan soal.

### 2. [konversi.go]

```go
package main

import "fmt"

func main() {

	var celcius float64

	fmt.Print("Masukkan Suhu: ")
	fmt.Scanln(&celcius)

	fmt.Println(celcius + 273)
}


```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel nama tipe data float64 lalu memberikan input dari user dengan hasil yang diperoleh dari rumus hasil input ditambah 273 yang menghasilkan output sesuai dengan soal.


### 3. [tukar.go]

```go
package main

import "fmt"

func main() {

	var x, y, z int

	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}


```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel x, y, z lalu memberikan input dari user yang menghasilkan output dari hasil rumus nilai x menghasilkan nilai z, nilai z menghasilkan nilai y, nilai y menghasilkan nilai x dan menghasilkan output sesuai dengan soal.



## Unguided

### 1. [konversihari.go]

```go
package main

import "fmt"

func main() {

	var jumlahhari int

	fmt.Scan(&jumlahhari)

	tahun := jumlahhari / 360
	sisa := jumlahhari % 360

	bulan := sisa / 30
	sisa = sisa % 30

	minggu := sisa / 7
	sisa = sisa % 7

	sisahari := sisa % 7

	fmt.Println(jumlahhari, "=", tahun, "tahun", bulan, "bulan", minggu, "minggu", sisahari, "hari")
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png


#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel nilai uang tipe data integer, memberikan ruang input nominal uang di fmt.Scan, lalu menuliskan rumus dengan kombinasi metode pembagian bilangan bulat dan sisa hasil bagi yang menghasilkan output sesuai dengan soal.

### 2. [konversisuhu.go]

```go
package main

import "fmt"

func main() {
	var celcius, reamur float64

	fmt.Scan(&celcius)

	reamur = (4.0 / 5.0) * celcius

	fmt.Println(reamur)
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/Kalkulator/output.png

#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel a,b serta tambah,kali,bagi,kurang, dan sisa hasil bagi tipe data integer, memberikan ruang input 2 nominal uang di fmt.Scan, lalu menuliskan rumus dengan metode tambah,kali,bagi,kurang, dan sisa hasil yang menghasilkan output sesuai dengan soal.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "cacahuang.go"
### 2. "kalkulator.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan dasar penulisan struktur format,penulisan variabel, print untuk input serta output, merangkai rumus dari hasil variabel dan input serta selain itu cara push dan commit ke github beserta running nya.

## Referensi
1. [ariefyusufw@telkomuniversity.ac.id]. ([2025]). *[Mengenal Pemrograman Golang: Konsep, Sintaks, dan Praktiknya!]*. Diakses pada [27 september 2026] melalui [https://it.telkomuniversity.ac.id/mengenal-pemrograman-golang-konsep-sintaks-dan-praktiknya/]

