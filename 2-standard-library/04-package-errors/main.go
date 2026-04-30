package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound    = errors.New("data tidak di temukan")
	ErrUnauhorized = errors.New("anda tidak memiliki akses")
	ErrOutOfStock  = errors.New("Stok habis")
)

func beliProduk(stok, qty int) error {
	if stok < qty {
		// wrap error
		return fmt.Errorf("gagal beli produk (stok: %d, qty: %d): %w", stok, qty, ErrOutOfStock)
	}
	return nil
}

func main() {
	wrapped := fmt.Errorf("level 1: %w", ErrNotFound)
	fmt.Println("Wrapped  :", wrapped)
	fmt.Println("Unwrapped:", errors.Unwrap(wrapped))

	//Latihan : Buat error sentinel ErrOutOfStock dan function beliProduk() yang mengembalikan wrapped error. Gunakan errors.Is() untuk mengeceknya.
	err := beliProduk(5, 10)

	if err != nil {
		fmt.Println("Error:", err)

		// cek pakai errors.Is
		if errors.Is(err, ErrOutOfStock) {
			fmt.Println("Handling khusus: stok tidak cukup")
		}
	}
}
