package main

import "fmt"

func somar(num1 float64, num2 float64) {

	resultado := num1 + num2

	imprimir(resultado)

}

func imprimir(num float64) {
	fmt.Println(num)
}

func main() {

	somar(10, 30)

}
