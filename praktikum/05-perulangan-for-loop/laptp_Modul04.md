# <h1 align="center">Tugas Pendahuluan Modul [05] - [Perulangan-for-loop]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

### 1. ganjil.go
```go
package main

import "fmt"


func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	
	for i := 1; i <= n; i += 2 {
		fmt.Printf("%d ", i)
		sum += i
	}
	fmt.Println()
	
	fmt.Println(sum)
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/ganjil/Screenshot%202026-10-15%20070935.png


#### Deskripsi
Pada soal kali ini diberikan inisialisasi nilai 1 lalu kondisi i lebih kecil atau sama dengan hasil input(n) lalu post statement dengan menambahkan digit dari angka awal(misal 1 akan ditambah 2 untuk membentuk angka ganjil) lalu dibawahnya ada fmt.Println dengan format verb dari %d untuk menyatakan variabel dari hasil i tersebut lalu ada sum untuk menjumlahkan nilai

### 2. perkalian.go

```go
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/05-perulangan-for-loop/tp/perkalian/Screenshot%202026-10-15%20071156.png


#### Deskripsi
pada kode kali ini diberikan 2 input a dan b, lalu user memberikan input a dan b tersebut, selanjutnya pakai inisialisasi 1 untuk nilai awal dan kondisi 1 lebih kecil atau sama dengan b dan melakukan increment untuk menambahkan nilai variabel pada iterasi, lalu print hasil i tadi menggunakan format verb %d * %d untuk menampung input i tadi(bukan untuk menuliskan rumus karena fungsi tanda "" untuk mengapit teks, lalu ada tanda \n setelah sudah ada hasil maka akan dipindah ke baris baru, lalu ada a masuk pada %d pertama dan i masuk pada %d kedua dan a * i masuk ke %d setelah =

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan penulisan for, inisialisasi, kondisi, dan post statement