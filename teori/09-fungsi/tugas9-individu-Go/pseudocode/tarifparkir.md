FUNCTION hitungTarifParki(lama_parkir)
    IF lama_parkir <= 2 THEN
        total_tarif = 5000
    ELSE
        jam_berikutnya = lama_parkir - 2
        total_tarif = 5000 + (jam_berikutnya * 3000)
    END IF
    
    RETURN total_tarif
END FUNCTION