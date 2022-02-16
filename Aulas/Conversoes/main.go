package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 1
	y := 2.4

	fmt.Println(float64(x) / y)

	fmt.Println("Cuidado para não converter string assim...", string(x))

	fmt.Println("Muito melhor...", strconv.Itoa(x))

}
