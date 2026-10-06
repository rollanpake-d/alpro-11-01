# <h1 align="center">Laporan Praktikum Modul [02] - [bahasa-pemrograman-go]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [Judul Topik Dasar Teori 1, Metode Kerja Git & Github]
Penjelasan: Topik ini berbicara mulai dari cara instalasi Git pada OS Windows, Pembuatan akun Github serta repositorinya, dan cara membuat struktur folder yang rapi dan benar, membuat kloning repositori dari dosen, konfigurasi Remote URL & identitas, mengubah berkas & menggunggah perubahan(push)

### B. [Judul Topik Dasar Teori 2, Bahasa Pemrograman Go]

#### 1. [Judul Sub-topik 1, Struktur Pemrograman Go]
Penjelasan: package main adalah penanda bahwa file ini berisi program utama, func main berisi kode utama sebuah program Go.


#### 2. [Judul Sub-topik 2, misal: Tipe Data dan Deklarasi Variabel di Go]
Penjelasan: notasi tipe ada integer, real, boolean, karakter, string. Deklarasi variabel di Go ada 3 yaitu deklarasi singkat(:=), tanpa deklarasi var, dan deklarasi banyak variabel(multiple variable)

## Guided

### 1. [lingkaran.go]

```go
package main

import "fmt"

func main() {
	var pi = 3.14
	var jarijarilingkaran, luaslingkaran float64

	//Membaca Input
	fmt.Scan(&jarijarilingkaran)

	//Menghitung Luas Lingkaran
	luaslingkaran = pi * jarijarilingkaran * jarijarilingkaran

	//Menampilkan Output
	fmt.Println(luaslingkaran)
}
```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel phi dan jari-jari lingkaran dalam tipe float, memberikan ruang input angka dengan fmt.Scan, lalu menuliskan rumus luas lingkaran dari variabel phi dikali jari-jari lingkaran sebanyak 2 kali dan menghasilkan output sesuai dengan soal.

### 2. [skor.go]

```go
package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca Input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//Menghitung Total dan Rata-Rata
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2

	//Menampilkan Output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}

```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel nama tipe data string serta skor matematika dan skor bahasa inggris tipe data integer, lalu memberikan input total dengan skor matematika ditambah skor bahasa inggris serta rata-rata dari hasil total dibagi 2 yang menghasilkan output sesuai dengan soal.


### 3. [suhu.go]

```go
package main

import "fmt"

func main() {
	var celcius float64
	var reamur, fahrenheit, kelvin float64

	//Membaca Input
	fmt.Scan(&celcius)

	//Menghitung Suhu
	reamur = celcius * 4.0 / 5.0
	fahrenheit = celcius*9.0/5.0 + 32.0
	kelvin = celcius + 273.15

	//Menampilkan output
	fmt.Println(reamur)
	fmt.Println(fahrenheit)
	fmt.Println(kelvin)

}

```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel celcius fahrenheit reamur serta kelvin, memberikan ruang input celcius dengan fmt.Scan, lalu menuliskan rumus menghitung suhu yang menghasilkan output sesuai dengan soal.

### 4. [tukar.go]

```go
package main

import "fmt"

func main() {
	var a, b int

	//Membaca Input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar Nilai
	a, b = b, a

	//Menampilkan Output
	fmt.Println(a)
	fmt.Println(b)
}

```
#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel a dan b, memberikan ruang input fmt.Scan, lalu menuliskan rumus variabel a atau b = variabel b atau a yang menghasilkan output sesuai dengan soal.

## Unguided

### 1. [cacahuang]

```go
package main

import "fmt"

func main() {
	var NominalUang int

	//Membaca Input
	fmt.Scan(&NominalUang)

	//Menghitung Cacah Uang
	Lembar10k := NominalUang / 10000
	sisa := NominalUang % 10000

	Lembar5k := sisa / 5000
	sisa = sisa % 5000

	Lembar1k := sisa / 1000

	//Menampilkan Output
	fmt.Println(Lembar10k, Lembar5k, Lembar1k)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png


#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel nilai uang tipe data integer, memberikan ruang input nominal uang di fmt.Scan, lalu menuliskan rumus dengan kombinasi metode pembagian bilangan bulat dan sisa hasil bagi yang menghasilkan output sesuai dengan soal.

### 2. [kalkulator]

```go
package main

import "fmt"

func main() {
	var a int
	var b int
	var tambah, kurang, kali, bagi, modulo int

	//Membaca Input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menghitung variabel
	tambah = a + b
	kurang = a - b
	kali = a * b
	bagi = a / b
	modulo = a % b

	//Menampilkan Output
	fmt.Println(tambah, kurang, kali, bagi, modulo)
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

