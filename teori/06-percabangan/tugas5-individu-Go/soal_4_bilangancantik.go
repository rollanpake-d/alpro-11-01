package main 

import "fmt"

func main () {
    var x int
    fmt.Scan(&x)

    cantik := true
    prevDigit := -1 

    for x > 0 {
        digit := x % 10
        
        if prevDigit != -1 && digit > prevDigit {
            cantik = false
            break 
        }
        
        prevDigit = digit
        x = x / 10
    }
    fmt.Println(cantik)
}