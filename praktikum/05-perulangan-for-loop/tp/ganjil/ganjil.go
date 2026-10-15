package main

import "fmt"


func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	
	for i := 1; i <= n; i += 2 {
		fmt.Printf("%d ", i)
		sum += i
	}
	fmt.Println()
	
	fmt.Println(sum)
}