package main

import (
	"fmt"
	"math"
)

func main() {

	a := 10
	b := 5

	soma := a + b
	fmt.Println("Soma =>", soma)

	subtracao := a - b
	fmt.Println("Subtração =>", subtracao)

	multiplicacao := a * b
	fmt.Println("Multiplicação =>", multiplicacao)

	divisao := a / b
	fmt.Println("Divisão =>", divisao)

	resto := a % b
	fmt.Println("Resto Divisão =>", resto)

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(float64(a), float64(b)))
	fmt.Println("Exponencial =>", math.Pow(float64(a), float64(b)))

}
