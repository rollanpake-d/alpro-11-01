# <h1 align="center">Laporan Praktikum Modul [06] - [perulangan-while-loop-dan-repeat]</h1>
<p align="center">[Paschalis Rolland Setianto Nugroho] - [109092600011]</p>

## Dasar Teori

### A. [While-loop]
Pada kondisi while loop, jika kondisi yang diinput bernilai salah dari ketentuan maka program akan berhenti, tetapi sebaliknya jika selama program benar maka program akan berjalan sampai batas pemberhentian yang udah ditentukan.

### B. [Repeat-Until]
Pada kondisi repeat until, jika kondisi yang diinput bernilai benar dari ketentuan maka program akan berhenti, tetapi sebaliknya jika selama program salah maka program akan berjalan sampai batas pemberhentian yang udah ditentukan.

## Guided

### 1. [terkecil.go]

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}

```
#### Deskripsi
Program Go di atas adalah program untuk mencari bilangan bulat positif terkecil yang nilai pangkat tiganya melebihi batas yang dimasukkan pengguna.

### 2. [terbesar.go]

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    }

    fmt.Println(n - 1)
}

```
#### Deskripsi
Program Go di atas digunakan untuk mencari bilangan bulat positif terbesar ($n-1$) yang kuadratnya ($n^2$) masih lebih kecil dari bilangan batas yang diinputkan pengguna.

### 3. [vokal.go]

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}

```
#### Deskripsi
Program ini menerima input sebuah kalimat dari pengguna, lalu menghitung dan menampilkan jumlah total huruf vokal (a, i, u, e, o) yang ada di dalam kalimat tersebut.


## Unguided

### 1. [atm.go]

```go
package main

import "fmt"

func main () {
	var password int
	fmt.Printf("Masukkan password ATM : ")

	max := 3
	input := 0
	saldo := 10000000

	for input < max {
		fmt.Scan(&password)

		if password == 1234 {
			fmt.Println("Login Berhasil!")
		
			pilihan := 0
				for pilihan != 3 {
				fmt.Println("Menu ATM:")
				fmt.Println("1. Lihat Saldo")
				fmt.Println("2. Tarik Uang")
				fmt.Println("3. Keluar")
				fmt.Println("Pilih: ") 

				fmt.Scan(&pilihan)

				if pilihan == 1 {
					fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
					var tarik int
					fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp")
					fmt.Scan(&tarik)

					if tarik <= saldo {
						saldo = saldo - tarik
						fmt.Printf("Anda berhasil menarik uang sebesar: Rp %d\n", tarik)
						fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
					} else {
					fmt.Println("Saldo tidak mencukupi")
					}
				} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
			}
		}
		break

		}	
		
		input++ 

		if input == max {
			fmt.Println("ATM terblokir. Program selesai. ")
		} else {
			fmt.Println("Password salah. Silakan coba lagi.")
			fmt.Printf("Masukkan password ATM : ")
		}
	}
}
```

##### Output
https://github.com/rollanpake-d/alpro-11-01/blob/589460344bcae5231d47eaba0aa7cf38318824c1/praktikum/06-perulangan-while-loop-dan-repeat/unguided/image.png

#### Deskripsi
Pada soal ini berfungsi untuk menyelesaikan persoalan sebuah proses login bank ketika berhasil akan dilanjutkan ke program melihat saldo, mengurangi saldo, dan keluar. disisi lain jika input password berbeda dari yang sudah ditentuin yaitu angka 1234 dan bila user selama 3 kali salah makan akan terblokir dan program tertutup/tidak bisa lanjut. dalam hal ini menggunakan while loop untuk proses selama benar maka program akan tetap bisa dijalankan dan akan langsung berhenti bila memenuhi kondisi yang sudah ditentukan 
<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->
### 1. "atm.go"

## Kesimpulan
Tujuan praktikum kali ini untuk mengenalkan aturan percabangan while loop, dan repeat until untuk membedakan dalam proses pemberhentian sebuah program itu, while loop program akan tetap berjalan selama kondisinya benar dan akan berhenti bila kondisinya salah, tapi jika repeat until itu ia akan berjalan bila kondisinya salah dan akan berhenti ketika benar.

## Referensi
1. Modul 06_Logika While & Repeat Until - sumber: Modul Informatics lab,Telkom University.