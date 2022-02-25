package main

import "fmt"

func somar(num1 float64, num2 float64) float64 {

	resultado := num1 + num2

	return resultado

}

func imprimir(num float64) {
	fmt.Println(num)
}

func doisRetornos() (float64, int) {
	return 2.3, 2
}

func main() {

	_, teste := doisRetornos()
	resultado := somar(10, 30)
	imprimir(resultado)
	imprimir(float64(teste))
}
