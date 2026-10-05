package main

import "fmt"

func main() {
	var n, m, i int

	fmt.Scan(&n, &m)

	for i = n; i <= m; i++ {
		fmt.Printf("%d", i)
	}
}
