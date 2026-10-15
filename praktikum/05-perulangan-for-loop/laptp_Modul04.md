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
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/evaluasiekspresikontrol/Screenshot%202026-10-08%20102647.png


#### Deskripsi
Pada soal Go kali ini akan dilakukan pengecekkan output dari soal dan ketentuang yang sudah diberikan pada tipe data boolean yang akan menghasilkan jawaban hanya true dan false tergantung dari fmt.Println yang kita inputkan.

### 2. Tracing: Evaluasi Pernyataan Kondisi Go

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
https://github.com/rollanpake-d/alpro-11-01/blob/main/praktikum/04-runtunan-sekuensi/tp/tracing/Screenshot%202026-10-08%20150304.png


#### Deskripsi
Pada soal Go kali ini sudah diberikan rangkaian code yang berfungsi untuk pengecekkan kondisi hasil melalui 4 blok if-else

### A. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi? 
hasilnya yaitu 25

### B. Apa output yang dihasilkan oleh program? 
hasilnya yaitu "Nilai akhir variabel result: 25"

### C-1. Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya? 
iya, Langkah selanjutnya yaitu Program masuk ke dalam blok if lalu melakukan evaluasi y < 10. Nah karena y itu 5, jadi 5 < 10 nilainya true.

### C-2. Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini memengaruhi nilai result? 
iya benar, Pengaruhnya di baris ke 20 result += z.

### C-3. Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang terjadi? 
sebenarnya ada 1 true dan 1 false nya, bagian x == 10 true dan y > 10 false tapi karena ini berbentuk || OR maka semua benar.

### C-4. Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Apa dampaknya pada result? 
sebenarnya jawaban pada tanda kurung dekat if bernilai true dan true tapi karena ada bentuk ! NOT maka nilainya true berubah jadi salah sehingga terpilih result -= 10

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
langkah awal yaitu membaca soal lalu menentukan variable tahun tipe data int dan bulan tipe data string, lalu memberikan input untuk user berupa tahun dan bulan, membagi bulan sesuai dengan kategori 31 hari dan kategori 30 hari, memberikan percabangan pada bagian khusus tahun kabisat bulan april dengan modulus, lalu memberikan default untuk memberikan cadangan bila data yang diinput diluar dari ketentuang yang sudah ditentukan maka akan diberikan output "-" dan tidak memunculkan hasil

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan penulisan if else percabangan ketika suatu kondisi tidak memenuhi ketentuan pertama maka akan dialihkan ke kondisi yang lainnya.