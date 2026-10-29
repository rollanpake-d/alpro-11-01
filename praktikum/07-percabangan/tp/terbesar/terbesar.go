package main 

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println(a)
	} else if b >= a && b >= c {
		fmt.Println(b)
	} else if c >= b && c >= a {
		fmt.Println(c)
	}
}