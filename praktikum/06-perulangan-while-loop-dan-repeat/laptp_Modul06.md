# <h1 align="center">Tugas Pendahuluan Modul [06] - [06-perulangan-while-loop-dan-repeat]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. lulus.go
```go
package main

import "fmt"

func main () {
	var nilai int
	var status string

	for nilai != -1 {
		fmt.Print("Masukkan nilai anda : ")
		fmt.Scan(&nilai)
	 
	if nilai == -1 {
		status = "Program selesai"
	} else if nilai < 60 {
		status = "Anda tidak lulus"
	} else if nilai >= 60 {
		status = "Anda lulus ujian"
	}
	fmt.Println(status)
}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/29a30471f149bb0f1e84dbb99349219e9d45cf37/praktikum/06-perulangan-while-loop-dan-repeat/tp/lulus/image.png

#### Deskripsi
Pada soal kali ini digunakan untuk memecahkan masalah seleksi status kelulusan, user memberikan input berupa variabel nilai jika diinputkan -1 program akan langsung selesai, jika diinputkan nilai lebih kecil dari 60 program akan memberikan status "anda tidak lulus", jika diinputkan nilai lebih besar atau sama dengan 60 program akan memberikan status "anda lulus ujian"

### 2. deret.go

```go
package main 

import "fmt"

func main () {
	var n int

	fmt.Scan(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/29a30471f149bb0f1e84dbb99349219e9d45cf37/praktikum/06-perulangan-while-loop-dan-repeat/tp/deret/image.png

#### Deskripsi
pada kode kali ini digunakan untuk menyelesaikan masalah penjumlahan dari sebuah deret bilangan tersebut, jika nilai memenuhi kondisi dimana inisialisasi lebih kecil atau sama dengan variabel yang diinputkan maka program akan berjalan dan menjumlahkan deret bilangan selama kondisi masih terpenuhi sesuai dengan hasil inputnya

### 3. duaterendah.go

```go
package main

import "fmt"

func main() {
	var n, nilai int
	var namaSiswa string
	var namaMin1, namaMin2 string

	min1 := 99999
	min2 := 99999

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&namaSiswa, &nilai)

		//bagian blok if dan else if saya ambil dari hasil AI
		if nilai < min1 {
			min2 = min1
			namaMin2 = namaMin1
			min1 = nilai
			namaMin1 = namaSiswa
		} else if nilai < min2 {
			min2 = nilai
			namaMin2 = namaSiswa
		}
	}
	fmt.Println(namaMin1)
	fmt.Println(namaMin2)
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