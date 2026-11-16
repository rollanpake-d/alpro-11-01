// Definisi fungsi untuk menentukan letak titik
FUNCTION cekLetakTitik(x, y)
    IF x == 0 AND y == 0 THEN
        RETURN "Titik Asal"
    ELSE IF y == 0 THEN
        RETURN "Sumbu X"
    ELSE IF x == 0 THEN
        RETURN "Sumbu Y"
    ELSE IF x > 0 AND y > 0 THEN
        RETURN "Kuadran I"
    ELSE IF x < 0 AND y > 0 THEN
        RETURN "Kuadran II"
    ELSE IF x < 0 AND y < 0 THEN
        RETURN "Kuadran III"
    ELSE
        RETURN "Kuadran IV"
    END IF
END FUNCTION //function ini dari generate AI

// Program Utama
LetakTitikKartesius
    // Membaca masukan koordinat x dan y
    BACA x, y
    
    // Memanggil fungsi dan menyimpan hasilnya
    hasil = cekLetakTitik(x, y)
    
    // Menampilkan keluaran
    PRINT hasil
END