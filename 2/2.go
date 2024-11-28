package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	// Inisialisasi array dengan kapasitas x untuk menampung berat ikan
	var berat [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	// Hitung banyaknya wadah
	jumlahWadah := (x + y - 1) / y

	// Hitung total berat ikan di setiap wadah
	var totalBerat []float64
	for i := 0; i < jumlahWadah; i++ {
		var sum float64
		for j := i * y; j < min(i*y+y, x); j++ {
			sum += berat[j]
		}
		totalBerat = append(totalBerat, sum)
	}

	// Hitung rata-rata berat ikan di setiap wadah
	var totalSemua float64
	for _, beratWadah := range totalBerat {
		totalSemua += beratWadah
	}
	rataRata := totalSemua / float64(jumlahWadah)

	// Cetak hasil
	fmt.Println(totalBerat)
	fmt.Printf("Rata-rata berat ikan di setiap wadah: %.2f\n", rataRata)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
