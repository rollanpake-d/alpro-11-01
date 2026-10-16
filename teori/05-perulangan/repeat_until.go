package main

import "fmt"

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)

    // Inisialisasi variabel iterasi
    i := 1
    for {
        fmt.Printf("%d ", i)
        i++
        // Kondisi berhenti (until)
        if i > N {
            break
        }
    }
    fmt.Println()
}