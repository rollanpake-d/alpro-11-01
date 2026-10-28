package main

import "fmt"

func main () {
	var password int
	fmt.Printf("Masukkan password ATM : ")

	max := 3
	input := 0
	saldo := 10000000

	for input < max {
		fmt.Scan(&password)

		if password == 1234 {
			fmt.Println("Login Berhasil!")
		
			pilihan := 0
				for pilihan != 3 {
				fmt.Println("Menu ATM:")
				fmt.Println("1. Lihat Saldo")
				fmt.Println("2. Tarik Uang")
				fmt.Println("3. Keluar")
				fmt.Println("Pilih: ") 

				fmt.Scan(&pilihan)

				if pilihan == 1 {
					fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
					var tarik int
					fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp")
					fmt.Scan(&tarik)

					if tarik <= saldo {
						saldo = saldo - tarik
						fmt.Printf("Anda berhasil menarik uang sebesar: Rp %d\n", tarik)
						fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
					} else {
					fmt.Println("Saldo tidak mencukupi")
					}
				} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
			}
		}
		break

		}	
		
		input++ 

		if input == max {
			fmt.Println("ATM terblokir. Program selesai. ")
		} else {
			fmt.Println("Password salah. Silakan coba lagi.")
			fmt.Printf("Masukkan password ATM : ")
		}
	}
}