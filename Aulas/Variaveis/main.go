package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * math.Pow(raio, 2)

	fmt.Println(area)

	const (
		a = 1
		b = 2
	)

	var (
		m = 1
	)

	var f, l bool = true, false

	fmt.Println(a, b, m, f, l)

}
