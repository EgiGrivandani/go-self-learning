package hitung

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("\n[SETUP] Persiapan sebelum semua test...")
	exitCode := m.Run()

	fmt.Println("[TEARDOWN] Pembersihan setelah semua test...")
	os.Exit(exitCode)
}

func setupTest(t *testing.T, nama string) func() {
	t.Helper()
	t.Logf("[BEFORE] %s dimulai", nama)
	return func() {
		t.Logf("[AFTER ] %s selesai", nama)
	}
}

func TestDenganSetup1(t *testing.T) {
	teardown := setupTest(t, "TestDenganSetup1")
	defer teardown()
	assert.Equal(t, 15, Tambah(10, 5))
}

func TestDenganSetup2(t *testing.T) {
	teardown := setupTest(t, "TestDenganSetup2")
	defer teardown()
	assert.Equal(t, 5, Kurang(10, 5))
}
