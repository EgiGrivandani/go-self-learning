package hitung

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambahAssert(t *testing.T) {
	assert.Equal(t, 3, Tambah(1, 2))
	assert.Equal(t, 10, Tambah(-5, 15))
}

func TestKurangAssert(t *testing.T) {
	assert.Equal(t, 3, Kurang(4, 1))
	assert.Equal(t, 10, Kurang(15, 5))
}

func TestHitungLuas(t *testing.T) {
	result := HitungLuas(2, 10)

	assert.Equal(t, float64(20), result)
	assert.Positive(t, result)
	assert.Greater(t, result, float64(0))
}
