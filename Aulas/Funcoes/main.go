package main

import "fmt"

//função dentro de uma variável
var subtracao = func(num1, num2 int) int {
	result := num1 - num2
	return result
}

//função passando duas variáveis e retornando uma
func somar(num1 float64, num2 float64) float64 {

	resultado := num1 + num2

	return resultado

}

//função passando uma variável e não retornando nada
func imprimir(num float64) {
	fmt.Println(num)
}

//função passando nada e retornando duas variáveis
func doisRetornos() (float64, int) {
	return 2.3, 2
}

//função passando duas variáveis e retornando duas variáveis
func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {

	_, teste := doisRetornos()
	resultado := somar(10, 30)
	imprimir(resultado)
	imprimir(float64(teste))

	x, y := trocar(2, 1)
	fmt.Println(x, y)

	testeSub := subtracao(10, 5)

	imprimir(float64(testeSub))
}
