package main

import "fmt"

func main() {
    var bulat int
    fmt.Scan(&bulat)

    bil := bulat
    
    digitSebelumnya := bil % 10
    bil = bil / 10

    status := ""

    for bil > 0 {
        digitSekarang := bil % 10

        if digitSekarang > digitSebelumnya {
            if status == "ascending" {
                status = "tidak terurut"
                break
            }
            status = "descending"
        } else if digitSekarang < digitSebelumnya {
            if status == "descending" {
                status = "tidak terurut"
                break
            }
            status = "ascending"
        }
        digitSebelumnya = digitSekarang
        bil = bil / 10
    }

    if status == "" {
        status = "ascending" 
    }
    fmt.Println(status)
}