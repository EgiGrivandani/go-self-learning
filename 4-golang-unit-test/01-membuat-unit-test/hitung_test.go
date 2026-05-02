package hitung

import (
	"testing"
)

func TestTambah(t *testing.T) {
	hasil := Tambah(10, 5)
	expected := 15
	if hasil != expected {
		t.Errorf("Tambah(10, 5) = %d; expected %d", hasil, expected)
	}
}

func TestKurang(t *testing.T) {
	hasil := Kurang(10, 5)
	expected := 5
	if hasil != expected {
		t.Errorf("Kurang(10, 5) = %d; expected %d", hasil, expected)
	}
}
