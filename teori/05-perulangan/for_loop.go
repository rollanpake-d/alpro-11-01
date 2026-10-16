package main

import "fmt"

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)

    // (inisialisasi, kondisi, increment)
    for i := 1; i <= N; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}