package main

import "fmt"

func multiplicacao(num1, num2 int) int {
	return num1 * num2
}

func exec(funcao func(int, int) int, p1, p2 int) int {

	return funcao(p1, p2)

}

func main() {
	resultado := exec(multiplicacao, 3, 8)
	fmt.Println(resultado)

}
