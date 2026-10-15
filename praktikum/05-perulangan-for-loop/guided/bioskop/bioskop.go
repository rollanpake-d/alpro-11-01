package main

import "fmt"

func main () {
	var t int
	fmt.Print("input t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalpendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisfilm, hari string
		var jumlahtiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisfilm, &hari, &jumlahtiket)

		hargasatuan := 0

		if jenisfilm == "Reguler" {
			if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
				hargasatuan = 40000
			} else {
				hargasatuan = 50000
			}

		} else if jenisfilm == "3D" {
			if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
				hargasatuan = 60000
			} else {
				hargasatuan = 70000
			}

		} else if jenisfilm == "IMAX" {
			if hari == "senin" || hari == "selasa" || hari == "rabu" || hari == "kamis" {
				hargasatuan = 80000
			} else {
				hargasatuan = 90000
			}
		}
			totalharga := hargasatuan * jumlahtiket

			if jumlahtiket >= 5 {
				diskon := (totalharga * 10) / 100
				totalharga = totalharga - diskon
			}

			fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalharga)

			totalpendapatan = totalpendapatan + totalharga
		}

		fmt.Println("output total pendapatan:", totalpendapatan)
	}
