package main

import "fmt"

func main () {
	var huruf string
	fmt.Scan(&huruf)

	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" ||
	   huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("bukan konsonan")
	   } else {
		fmt.Println("konsonan")
	   }
}