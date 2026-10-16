package main

import "fmt"

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)
	
    i := 1
    // dengan kondisi
    for i <= N {
        fmt.Printf("%d ", i)
        i++ // Increment nilai i biar perulangan berhenti
    }
    fmt.Println()
}