package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Println("No 1 :", intNum > 5)
	fmt.Println("No 2 :", intNum >= 5 && intOther < 11)
	fmt.Println("No 3 :", sngNum != -1 || intOther < 0)
	fmt.Println("No 4 :", !(intNum > 3) || intNum <= 5)
	fmt.Println("No 5 :", !(intOther >= intNum))
	fmt.Println("No 6 :", 0-sngNum > 0)
	fmt.Println("No 7 :", 4/2 == intOther/intNum)
	fmt.Println("No 8 :", intOther%2 == 0)
	fmt.Println("No 9 :", intOther+2*intNum != 30 || !(sngNum > 0))
	fmt.Println("No 10:", intOther > 0 && intNum > 0 || sngNum > 0)
	fmt.Println("No 11:", sngNum > 0 || (intNum >= 0 && -1*intOther == -10))
	fmt.Println("No 12:", intNum == 5)
	fmt.Println("No 13:", intNum > 0 || (sngNum <= 0 && intOther == 13))
	fmt.Println("No 14:", !(!(!(!(intNum > 0)))))
}