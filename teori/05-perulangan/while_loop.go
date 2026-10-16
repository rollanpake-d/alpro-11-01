package main

import "fmt"

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)
	
    i := 1
    // Struktur while-loop di Go menggunakan keyword 'for' dengan kondisi
    for i <= N {
        fmt.Printf("%d ", i)
        i++ // Increment nilai i agar perulangan bisa berhenti
    }
    fmt.Println()
}