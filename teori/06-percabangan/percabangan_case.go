package main

import "fmt"

func main () {
	lamp_status := "kuning"

	if lamp_status == "merah" {
		fmt.Println("berhenti")
	} else if lamp_status == "kuning" {
		fmt.Println("hati-hati / lebih pelan")
	} else if lamp_status == "hijau" {
		fmt.Println("jalan")
	} else {
		fmt.Println("perintah tidak terdaftar disini")
	}
}