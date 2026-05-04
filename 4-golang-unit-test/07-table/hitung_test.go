package hitung

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambahTable(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"PositifDanPositif", 10, 5, 15},
		{"NegatifDanPositif", -10, 5, -5},
		{"NegatifDanNegatif", -10, -5, -15},
		{"NolDanPositif", 0, 5, 5},
		{"NolDanNol", 0, 0, 0},
		{"BesarDanBesar", 1000, 2000, 3000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Tambah(tc.a, tc.b))
		})
	}
}

func TestBagitTable(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		wantError bool
	}{
		{"Normal", 10, 2, 5.0, false},
		{"Desimal", 10, 3, 3.333, false},
		{"DenganNol", 10, 0, 0, true},
		{"Negatif", -10, 2, -5.0, false},
		{"SamaDenganDiriSendiri", 7, 7, 1.0, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hasil, err := Bagi(tc.a, tc.b)
			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.InDelta(t, tc.expected, hasil, 0.01)
			}
		})
	}
}
