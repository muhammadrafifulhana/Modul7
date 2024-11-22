package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	titikPusat Titik
	radius     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func diDalam(c Lingkaran, p Titik) bool {
	return jarak(c.titikPusat, p) <= float64(c.radius)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var px, py int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&px, &py)

	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{px, py}

	if diDalam(lingkaran1, titik) && diDalam(lingkaran2, titik) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam(lingkaran1, titik) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam(lingkaran2, titik) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
