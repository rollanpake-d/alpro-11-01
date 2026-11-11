package main

import "fmt"

func cekPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func memenuhiKriteria(n int) bool {
	temp := n
	jumlahDigit := 0
	digitSebelumnya := 10

	for temp > 0 {
		digit := temp % 10
		if digit > digitSebelumnya {
			return false 
		}
		digitSebelumnya = digit
		jumlahDigit += digit
		temp /= 10
	}
	return cekPrima(jumlahDigit)
}

func main() {
	var batasBawah, batasAtas int
	if _, err := fmt.Scan(&batasBawah, &batasAtas); err != nil {
		return
	}

	totalValid := 0
	for i := batasBawah; i <= batasAtas; i++ {
		if memenuhiKriteria(i) {
			totalValid++
		}
	}
	fmt.Println(totalValid)
}