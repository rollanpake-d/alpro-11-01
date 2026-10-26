package main

import "fmt"

func main () {
	var a, b ,c ,d, max, min int

	fmt.Scan(&a, &b ,&c ,&d)
	max = a
	min = a

	if b > max {
    max = b
	}
	if c > max {
    max = c
	}
	if d > max {
    max = d
	}

	if b < min {
    min = b
	}
	if c < min {
    min = c
	}
	if d < min {
    min = d
	}
	fmt.Println(max, min)
}