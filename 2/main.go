package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Printf("Masukkan %d elemen array:\n", n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("a. Tampilkan seluruh isi array")
		fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen dengan indeks genap")
		fmt.Println("d. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("e. Hapus elemen pada indeks tertentu")
		fmt.Println("f. Hitung rata-rata")
		fmt.Println("g. Hitung standar deviasi")
		fmt.Println("h. Hitung frekuensi bilangan tertentu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan string
		fmt.Scan(&pilihan)

		switch pilihan {
		case "a":
			fmt.Println("Seluruh isi array:", arr)
		case "b":
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < len(arr); i += 2 {
				fmt.Printf("%d ", arr[i])
			}
			fmt.Println()
		case "c":
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < len(arr); i += 2 {
				fmt.Printf("%d ", arr[i])
			}
			fmt.Println()
		case "d":
			var x int
			fmt.Print("Masukkan bilangan x untuk kelipatan indeks: ")
			fmt.Scan(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			for i := 0; i < len(arr); i++ {
				if i%x == 0 {
					fmt.Printf("%d ", arr[i])
				}
			}
			fmt.Println()
		case "e":
			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < len(arr) {
				arr = append(arr[:idx], arr[idx+1:]...)
				fmt.Println("Isi array setelah penghapusan:", arr)
			} else {
				fmt.Println("Indeks tidak valid!")
			}
		case "f":
			sum := 0
			for _, val := range arr {
				sum += val
			}
			rata := float64(sum) / float64(len(arr))
			fmt.Printf("Rata-rata: %.2f\n", rata)
		case "g":
			sum := 0
			for _, val := range arr {
				sum += val
			}
			rata := float64(sum) / float64(len(arr))

			var variance float64
			for _, val := range arr {
				variance += math.Pow(float64(val)-rata, 2)
			}
			variance /= float64(len(arr))
			stdDev := math.Sqrt(variance)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)
		case "h":
			var target int
			fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
			fmt.Scan(&target)
			count := 0
			for _, val := range arr {
				if val == target {
					count++
				}
			}
			fmt.Printf("Frekuensi %d: %d kali\n", target, count)
		case "0":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
