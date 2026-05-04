package hitung

import "testing"

func TestTambah(t *testing.T) {
	hasil := Tambah(2, 2)
	if hasil != 4 {
		t.Errorf("Tambah(2,2) = %d; expected 4", hasil)
	}
	t.Log("\nBaris ini TETAP jalan setelah t.Error")
}
func TestDenganFatal(t *testing.T) {
	// t.Fatal: catat error lalu TEST BERHENTI
	_, err := Bagi(10, 0)
	if err == nil {
		t.Fatal("Seharusnya error! Test berhenti di sini.")
	}
	t.Log("Baris ini TIDAK jalan jika Fatal() dipanggil di atas")
}
