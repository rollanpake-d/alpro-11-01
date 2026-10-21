package main

import "fmt"

func main() {
    var k int
    fmt.Scan(&k)

    var total int

    for i := 0; i < 5; i++ {
        var v int
        fmt.Scan(&v)

        total += v
        fmt.Println(total >= k)
    }
}