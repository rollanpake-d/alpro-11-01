package main

import "fmt"

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

// rataRata mengembalikan rata-rata skor dari n player
func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

// idxTertinggi mengembalikan indeks player dengan skor tertinggi;
// jika ada skor sama, ambil yang pertama muncul
func idxTertinggi(T tabPlayer, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idxMax].score {
			idxMax = i
		}
	}
	return idxMax
}

// hitungDiAtas mengembalikan banyaknya player dengan skor lebih dari batas
func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabPlayer

	fmt.Scan(&n)
	if n < 1 {
		n = 1
	} else if n > NMAX {
		n = NMAX
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	avg := rataRata(T, n)
	bestIdx := idxTertinggi(T, n)
	aboveCount := hitungDiAtas(T, n, avg)

	fmt.Printf("Rata-rata skor: %.2f\n", avg)
	fmt.Printf("Player terbaik: %s (%d)\n", T[bestIdx].nickname, T[bestIdx].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", aboveCount)
}