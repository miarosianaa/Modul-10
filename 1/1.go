package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var berat [1000]float64

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	minBerat := berat[0]
	maxBerat := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}

	// Mencetak hasil
	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
