// Definisi fungsi untuk konversi ke sentimeter
FUNCTION konversi(km, m, cm)
    // 1 km = 100.000 cm, 1 m = 100 cm
    total_cm = (km * 100000) + (m * 100) + cm
    RETURN total_cm
END FUNCTION

// Program Utama
KonversiJarak
    // Membaca masukan dari pengguna
    READ km, m, cm
    
    // Memanggil fungsi konversi
    hasil = konversi(km, m, cm)
    
    // Menampilkan keluaran
    PRINT "Hasil konversi = ", hasil, " cm"
END