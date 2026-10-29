# <h1 align="center">Tugas Pendahuluan Modul [7] - [07-percabangan]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. bilangan.go
```go
package main

import "fmt"

func main () {
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan % 2 == 0 && bilangan > 0 {
		fmt.Println("Genap, Positif")
	} else if bilangan % 2 == 0 && bilangan < 0 {
		fmt.Println("Genap, Negatif")
	} else if bilangan % 2 != 0 && bilangan > 0 {
		fmt.Println("Ganjil, Positif")
	} else if bilangan % 2 != 0 && bilangan < 0 {
		fmt.Println("Ganjil, Negatif")
	} else {
		fmt.Println("Nol")
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/29a30471f149bb0f1e84dbb99349219e9d45cf37/praktikum/06-perulangan-while-loop-dan-repeat/tp/lulus/image.png

#### Deskripsi
Pada soal kali ini digunakan untuk memecahkan masalah seleksi status kelulusan, user memberikan input berupa variabel nilai jika diinputkan -1 program akan langsung selesai, jika diinputkan nilai lebih kecil dari 60 program akan memberikan status "anda tidak lulus", jika diinputkan nilai lebih besar atau sama dengan 60 program akan memberikan status "anda lulus ujian"

### 2. hari.go

```go
package main 

import "fmt"

func main () {
	var nomorHari string
	
	fmt.Scan(&nomorHari)

	switch nomorHari {
	case "1": {
		fmt.Println("Senin")
	}
	case "2": {
		fmt.Println("Selasa")
	}
	case "3": {
		fmt.Println("Rabu")
	}
	case "4": {
		fmt.Println("Kamis")
	}
	case "5": {
		fmt.Println("Jumat")
	}
	case "6": {
		fmt.Println("Sabtu")
	}
	case "7": {
		fmt.Println("Minggu")
	}
	default:
		fmt.Println("Nomor hari tidak valid ")
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/29a30471f149bb0f1e84dbb99349219e9d45cf37/praktikum/06-perulangan-while-loop-dan-repeat/tp/deret/image.png

#### Deskripsi
pada kode kali ini digunakan untuk menyelesaikan masalah penjumlahan dari sebuah deret bilangan tersebut, jika nilai memenuhi kondisi dimana inisialisasi lebih kecil atau sama dengan variabel yang diinputkan maka program akan berjalan dan menjumlahkan deret bilangan selama kondisi masih terpenuhi sesuai dengan hasil inputnya

### 3. terbesar.go

```go
package main 

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println(a)
	} else if b >= a && b >= c {
		fmt.Println(b)
	} else if c >= b && c >= a {
		fmt.Println(c)
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/29a30471f149bb0f1e84dbb99349219e9d45cf37/praktikum/06-perulangan-while-loop-dan-repeat/tp/duaterendah/image.png


#### Deskripsi
pada kode kali ini digunakan untuk mencari variabel nama dengan status nilai terendah diantara variabel lainnya, ditentukan dahulu untuk 2 nama siswa melalui min1 dan min2, fungsinya untuk memberikan identitas khusus 2 nama dengan nilai terendah lalu diberikan sebuah input jumlah siswa yang akan dimasukkan, misal 5 maka hanya bisa menginputkan data 5 siswanya lalu input nama siswanya juga sama nilainya. saat variabel min1 nilainya lebih besar dari min2 nya maka akan digeser ke min2, kemudian karena min1 sudah kosong maka diganti dengan data yang terbaru, lalu bagian else if jika nilai min 1 lebih kecil dari min2 maka min1 tidak diubah karena nilai min2nya masih lebih besar, nah jika tidak keduanya maka data baru tidak dianggap karena nilainya tidak masuk kategori 2 terendahnya

### 4. biaya.go

```go
package main

import "fmt"

func main () {
	var n int
	var biayaKuliah float64
	
fmt.Scan(&n)

biayaKuliah = 10000000

	for i := 2; i <= n; i++ {
		biayaKuliah = biayaKuliah * 1.06
	}
	fmt.Printf("%.0f", biayaKuliah)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/29a30471f149bb0f1e84dbb99349219e9d45cf37/praktikum/06-perulangan-while-loop-dan-repeat/tp/biaya/image.png


#### Deskripsi
pada kode kali ini digunakan untuk mengetahui hasil penjumlahan biaya kuliah yang mengalami kenaikan sebesar 0,6 persen per tahunnya dari biaya awal kuliah sebesar 10.000.000 juta

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan kondisi while loop, dimana kode akan tetap dijalankan ketika masih memenuhi kondisi(masih true), dan repeat until, dimana program akan dihentikan ketika ada nilai yang tidak memenuhi syarat(hingga false)