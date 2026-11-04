package main

import "fmt"

func main() {
	var t, b, th int
	fmt.Scan(&t, &b, &th)

	kabisat := (th%4 == 0 && th%100 != 0) || (th%400 == 0)

	hari := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if kabisat {
		hari[2] = 29
	}

	//hasil generate AI
	if b < 1 || b > 12 || t < 1 || t > hari[b] {
		fmt.Println("Tidak valid")
		return
	}

	fmt.Println("Valid")
	
	t, b, th = t+1, b, th
	if t > hari[b] {
		t, b = 1, b+1
		if b > 12 {
			b, th = 1, th+1
		}
	}
	fmt.Printf("%d/%d/%d\n", t, b, th)
}