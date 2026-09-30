# <h1 align="center">Tugas Pendahuluan Modul [03] - [Tipe data dan instruksi data]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. Sisa.go

```go
package main

import "fmt"

func main() {

	var y, x int

	//input
	fmt.Scan(&y, &x)

	//output
	fmt.Println(y % x)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/Screenshot%202026-09-30%20153655.png


#### Deskripsi
langkah awal yaitu membaca soal lalu menentukan variable menggunakan huruf y dan x dengan tipe data integer, membuat input untuk user dari untuk menghasilkan variable y dan x, buat print dari hasil y modulo x dan menghasilkan input sesuai pada soal di modul

### 2. bool.go

```go
package main

import "fmt"

func main() {
	var inputBool bool

	fmt.Scan(&inputBool)

	fmt.Println(inputBool)
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/Screenshot%202026-09-30%20153746.png


#### Deskripsi
langkah awal yaitu membaca soal lalu menentukan variable menggunakan inputbool bool, membuat input untuk user dari variabel inputbool, buat print dari hasil inputbool dan menghasilkan input sesuai pada soal di modul.

### 2. konversi.go

```go
package main

import "fmt"

func main() {
	var mil float64

	fmt.Scan(&mil)

	km := mil * 1.6

	fmt.Printf("%.1f\n", km)
}

```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi/Screenshot%202026-09-30%20153831.png

#### Deskripsi
langkah awal yaitu membaca soal lalu menentukan variable mil tipe data float64, membuat input untuk user dari variabel mil, buat print dari hasil rumus input dikali 1.6 (1 km) dan menghasilkan input sesuai pada soal di modul.

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan dasar penulisan struktur format,penulisan variabel, print untuk input serta output, merangkai rumus dari hasil variabel dan input serta selain itu cara push dan commit ke github beserta running nya.