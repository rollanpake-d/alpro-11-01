package main 

import "fmt"

func main () {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)

	donat := 0

	for x <= n && y <= m {
		n -= x 
		m -= y
		donat++
	}
	fmt.Print(donat)
}