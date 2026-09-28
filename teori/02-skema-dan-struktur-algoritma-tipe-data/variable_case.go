package main

import "fmt"

func main() {

	var name string
	name = "Paschalis Rolland Setianto Nugroho"

	var middlename = "Paschalis Rolland Setianto Nugroho"
	fmt.Println("nama depan", middlename)

	var lastname = "Paschalis Rolland Setianto Nugroho"
	fmt.Println("nama depan", lastname)

	fmt.Println("nama : ", name)

	var (
		fullname  = "Paschalis Rolland"
		firstname = "Paschalis"
	)

	fmt.Println(fullname)
	fmt.Println(firstname)
}
