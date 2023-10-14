package main

import (
	"fmt"
)

func main() {
	// Nilai tagihan
	tagihan := 275

	// Hitung tip berdasarkan aturan yang diberikan
	var tip float64
	if tagihan >= 50 && tagihan <= 300 {
		tip = float64(tagihan) * 0.15
	} else {
		tip = float64(tagihan) * 0.20
	}

	// Hitung total nilai (tagihan + tip)
	total := float64(tagihan) + tip

	// Tampilkan hasil di konsol
	fmt.Printf("Tagihannya %d, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)
}
