### 1. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing. 

jika for dengan ketiga komponen(inisialisasi, kondisi, dan post statement) digunakan ketika nilai variabel sudah diketahui pada 1 blok kode, tapi jika hanya pakai kondisi maka digunakan ketika nilainya belum diketahui diluar daripada blok kode tersebut(bukan diblok yang melainkan sebelum blok itu).


### Contoh ketika pakai 3 komponen
```go
package main​

import "fmt"​

func main(){​
    var n, m, i int​

    fmt.Scan(&n,&m)​

    for i = n; i <= m; i++ {​

    fmt.Printf("%d", i)​
    } ​
}​
```

### Contoh ketika hanya kondisi

```go
package main

import "fmt"

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    i := n 

    for i <= m {
    fmt.Printf("%d ", i)
    i++ 
    }
}
```