package main

import "fmt"

// Fungsi untuk mengecek tahun kabisat
func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}

// Fungsi untuk mengembalikan jumlah hari pada bulan tertentu dalam suatu tahun
func jumlahHari(bulan, tahun int) int {
	var hari int
	if bulan == 2 {
		if kabisat(tahun) {
			hari = 29
		} else {
			hari = 28
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		hari = 30
	} else {
		hari = 31
	}
	return hari
}

// Fungsi untuk menghitung hari ke-berapa dalam tahun tersebut menggunakan perulangan
func hariKe(tanggal, bulan, tahun int) int {
	var total int = 0
	var i int
	for i = 1; i < bulan; i++ {
		total = total + jumlahHari(i, tahun)
	}
	total = total + tanggal
	return total
}

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	// Menghitung urutan hari dalam tahun
	urutanHari := hariKe(tanggal, bulan, tahun)

	// Menentukan jumlah total hari dalam satu tahun (kabisat = 366, biasa = 365)
	var totalHariTahun int
	if kabisat(tahun) {
		totalHariTahun = 366
	} else {
		totalHariTahun = 365
	}

	// Menghitung sisa hari hingga akhir tahun
	sisaHari := totalHariTahun - urutanHari

	// Menampilkan keluaran sesuai format
	fmt.Printf("Hari ke-%d\n", urutanHari)
	fmt.Printf("Sisa hari dalam tahun %d: %d\n", tahun, sisaHari)
}