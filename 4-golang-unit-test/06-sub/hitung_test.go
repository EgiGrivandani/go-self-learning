package hitung

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambahSub(t *testing.T) {
	t.Run("Positif", func(t *testing.T) {
		assert.Equal(t, 20, Tambah(4, 16))
		assert.Equal(t, 5, Tambah(2, 3))
	})

	t.Run("Negatif", func(t *testing.T) {
		assert.Equal(t, 2, Kurang(16, 14))
		assert.Equal(t, 5, Kurang(10, 5))
	})
}
