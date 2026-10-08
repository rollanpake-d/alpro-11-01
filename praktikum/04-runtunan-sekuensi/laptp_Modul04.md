# <h1 align="center">Tugas Pendahuluan Modul [04] - [Runtunan dan Sekuensi]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go 
```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Println("No 1 :", intNum > 5)
	fmt.Println("No 2 :", intNum >= 5 && intOther < 11)
	fmt.Println("No 3 :", sngNum != -1 || intOther < 0)
	fmt.Println("No 4 :", !(intNum > 3) || intNum <= 5)
	fmt.Println("No 5 :", !(intOther >= intNum))
	fmt.Println("No 6 :", 0-sngNum > 0)
	fmt.Println("No 7 :", 4/2 == intOther/intNum)
	fmt.Println("No 8 :", intOther%2 == 0)
	fmt.Println("No 9 :", intOther+2*intNum != 30 || !(sngNum > 0))
	fmt.Println("No 10:", intOther > 0 && intNum > 0 || sngNum > 0)
	fmt.Println("No 11:", sngNum > 0 || (intNum >= 0 && -1*intOther == -10))
	fmt.Println("No 12:", intNum == 5)
	fmt.Println("No 13:", intNum > 0 || (sngNum <= 0 && intOther == 13))
	fmt.Println("No 14:", !(!(!(!(intNum > 0)))))
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/evaluasiekspresikontrol/Screenshot%202026-10-08%20102647.png


#### Deskripsi
langkah awal yaitu membaca soal lalu menentukan variable menggunakan huruf y dan x dengan tipe data integer, membuat input untuk user dari untuk menghasilkan variable y dan x, buat print dari hasil y modulo x dan menghasilkan input sesuai pada soal di modul

### 2. Tracing: Evaluasi Pernyataan Kondisi Go

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/tracing/Screenshot%202026-10-08%20150304.png


#### Deskripsi
langkah awal yaitu membaca soal lalu menentukan variable menggunakan inputbool bool, membuat input untuk user dari variabel inputbool, buat print dari hasil inputbool dan menghasilkan input sesuai pada soal di modul.

### 3. Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan Go

```go
package main

import (
	"fmt"
)

func Kabisat(tahun int) bool {
	return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if Kabisat(tahun) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		fmt.Println("-")
	}
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/jumlah%20hari/Screenshot%202026-10-08%20145010.png

#### Deskripsi
langkah awal yaitu membaca soal lalu menentukan variable mil tipe data float64, membuat input untuk user dari variabel mil, buat print dari hasil rumus input dikali 1.6 (1 km) dan menghasilkan input sesuai pada soal di modul.

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan dasar penulisan struktur format,penulisan variabel, print untuk input serta output, merangkai rumus dari hasil variabel dan input serta selain itu cara push dan commit ke github beserta running nya.