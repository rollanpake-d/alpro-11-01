FUNCTION jumlahDeretKuadrat(N)
    total = 0
    UNTUK i DARI 1 SAMPAI N LAKUKAN total = total + (i * i)
    SELESAI UNTUK KEMBALIKAN total
END FUNCTION

// Program Utama
BACA N
CETAK jumlahDeretKuadrat(N)