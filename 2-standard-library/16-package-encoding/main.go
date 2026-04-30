package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Produk struct {
	Name   string  `json:"name"`
	Harga  int     `json:"harga"`
	Stok   int     `json:"stok"`
	Diskon float64 `json:"diskon,omitempty"` // omitempty: hilangkan jika kosong
}

func main() {
	str := "EgiGrivandani"
	encode := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("Encoded : ", encode)

	decoded, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Decode : ", string(decoded))
	}

	produk1 := Produk{
		Name:   "Sabun",
		Harga:  20000,
		Stok:   20,
		Diskon: 0.3,
	}

	jsonBytes, _ := json.Marshal(produk1)
	fmt.Println("JSON:", string(jsonBytes))

	// MarshalIndent - JSON yang rapi
	jsonPretty, _ := json.MarshalIndent(produk1, "", "  ")
	fmt.Println("Pretty:\n" + string(jsonPretty))
}
