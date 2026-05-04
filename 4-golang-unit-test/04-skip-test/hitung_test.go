package hitung

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipSederhana(t *testing.T) {
	t.Skip("skip TestSkipSederhana")
	assert.Equal(t, 2, Tambah(2, 1))
}

func TestSkipOs(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Hanya berjalan di Windows")
	}
	t.Log("Berjalan di Windows ✅")
	assert.Equal(t, 10, Tambah(5, 5))
}

func TestSkipBerdasarkanShortMode(t *testing.T) {
	if testing.Short() {
		t.Skip("Dilewati dalam mode -short")
	}
	t.Log("Test lengkap (bukan short mode)")
	assert.Equal(t, 1024, Pangkat(2, 10))
}
