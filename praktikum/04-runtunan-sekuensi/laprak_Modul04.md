# <h1 align="center">Laporan Praktikum Modul [04] - [runtunan-sekuensi]</h1>
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

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
	    grade = "D"
	} else {
		grade = "F"
	}
	
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

}
```
#### Deskripsi
Pada tugas guided ini diberikan soal yang berfungsi untuk memberikan hasil nama serta grade berdasarkan input yang diberikan dari user, menggunakan else if serta package bufio dan os yang baru pernah digunakan pada tugas kali ini.

### 2. [klasifikasi.go]

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}


```
#### Deskripsi
Pada tugas guided kali ini diberikan soal untuk memberikan output keterangan berdasarkan klasifikasi dari input usia dan gaji dengan menggunakan switch case, dimana hal ini bisa dikatakan lebih simple dari model else if karena switch case tidak perlu penambahan tanda kurung "()"


### 3. [penilaian.go]

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
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
