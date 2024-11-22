package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	fmt.Println("\nMasukkan skor pertandingan (masukkan skor negatif untuk berhenti):")

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: %s", pertandingan, klubA))
		} else if skorB > skorA {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: %s", pertandingan, klubB))
		} else {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: Draw", pertandingan))
		}

		pertandingan++
	}

	for _, h := range hasil {
		fmt.Println(h)
	}

	fmt.Println("Pertandingan selesai")
}
