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
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel x dalam tipe data integer, memberikan input variabel dengan fmt.Scan dari user, lalu menuliskan rumus pecahan uang menggunakan pembagian bilangan bulat serta modulus dan menghasilkan output sesuai dengan soal.

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
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel nama tipe data float64 lalu memberikan input dari user dengan hasil yang diperoleh dari rumus hasil input ditambah 273 yang menghasilkan output sesuai dengan soal.


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
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel x, y, z lalu memberikan input dari user yang menghasilkan output dari hasil rumus nilai x menghasilkan nilai z, nilai z menghasilkan nilai y, nilai y menghasilkan nilai x dan menghasilkan output sesuai dengan soal.



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
Proses awal mulai dari mengamati perintah pada soal lalu diawal menentukan variabel jumlahhari dengan tipe data integer lalu memberikan input dari user pada fmt.Scan, lalu memasukkan rumus mulai dari urutan tahun menggunakan ketentuan hari pada tahun lalu bulan, minggu dan sisahari dengan operator pembagian bulan (/) serta modulus (%) yang menghasilkan output sesuai dengan soal.

### 2. [pajak.go]

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
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/unguided/pajak/Screenshot%202026-10-11%20181946.png

#### Deskripsi
Proses awal mulai dari mengamati perintah pada soal lalu memberikan variabel celcius dan reamur dengan tipe data float64 ,lalu memberikan input yang akan diberikan dari user lalu menuliskan rumus variabel reamur = (4 per 5) dikali input dari celciusnya ,yang menghasilkan output sesuai dengan soal.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "konversihari.go"
### 2. "konversisuhu.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan dasar struktur penulisan variabel, tipe data yang digunakan serta kategori tipe datanya.

## Referensi
1. [Tipe & variabel] melalui https://ajarka.com/belajar/go/tipe-data

