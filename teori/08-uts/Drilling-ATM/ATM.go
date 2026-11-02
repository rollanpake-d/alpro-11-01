package main

import "fmt"

func main() {
	var pin string
	var percobaan int
	var saldo int = 10000000
	var pilihan int
	var jumlah int
	var login bool = false
	var mutasi string

	// Proses login ATM
	for percobaan < 3 && !login {
		fmt.Print("Masukkan PIN ATM: ")
		fmt.Scan(&pin)

		if pin == "1234" {
			login = true
			fmt.Println("Login berhasil!")
		} else {
			percobaan++
			fmt.Println("PIN salah. Silakan coba lagi.")
		}
	}

	// Jika PIN salah sebanyak 3 kali
	if !login {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	// Menu ATM
	for pilihan != 6 {
		fmt.Println("\n===== MENU ATM =====")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Setor Tunai")
		fmt.Println("4. Transfer Uang")
		fmt.Println("5. Mutasi Rekening")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)

		case 2:
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah penarikan harus lebih dari 0.")
			} else if jumlah <= saldo {
				saldo -= jumlah
				fmt.Printf("Berhasil menarik uang sebesar Rp %d\n", jumlah)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
				mutasi += fmt.Sprintf("\nTarik tunai: Rp %d", jumlah)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}

		case 3:
			fmt.Print("Masukkan jumlah uang yang ingin disetor: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah setoran harus lebih dari 0.")
			} else {
				saldo += jumlah
				fmt.Printf("Setor tunai berhasil sebesar Rp %d\n", jumlah)
				fmt.Printf("Saldo sekarang: Rp %d\n", saldo)
				mutasi += fmt.Sprintf("\nSetor tunai: Rp %d", jumlah)
			}

		case 4:
			var rekening string

			fmt.Print("Masukkan nomor rekening tujuan: ")
			fmt.Scan(&rekening)

			fmt.Print("Masukkan jumlah uang yang ingin ditransfer: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah transfer harus lebih dari 0.")
			} else if jumlah <= saldo {
				saldo -= jumlah
				fmt.Println("Transfer berhasil!")
				fmt.Println("Nomor rekening tujuan:", rekening)
				fmt.Printf("Jumlah transfer: Rp %d\n", jumlah)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
				mutasi += fmt.Sprintf("\nTransfer ke %s: Rp %d", rekening, jumlah)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}

		case 5:
			fmt.Println("\n===== MUTASI REKENING =====")
			fmt.Println(mutasi)
			fmt.Printf("Saldo akhir: Rp %d\n", saldo)
			fmt.Println("\n===========================")

		case 6:
			fmt.Println("Keluar dari sistem. Terima kasih!")

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}