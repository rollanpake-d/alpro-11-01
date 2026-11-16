FUNCTION faktorial(n)
    JIKA n <= 1 MAKA KEMBALIKAN 1 SELESAI JIKA hasil = 1
    UNTUK i DARI 2 SAMPAI n LAKUKAN hasil = hasil * i SELESAI UNTUK KEMBALIKAN hasil
END FUNCTION

FUNCTION hitungKombinasi(x, y)
    n = MAKS(x, y)
    r = MIN(x, y)
    KEMBALIKAN faktorial(n) / (faktorial(r) * faktorial(n - r))
END FUNCTION

// Program Utama
BACA x, y
CETAK faktorial(x), faktorial(y), hitungKombinasi(x, y)