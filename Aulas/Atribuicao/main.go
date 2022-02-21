package main

import (
	"fmt"
	"time"
)

func main() {

	var a = 3
	a += 2
	a -= 3
	a /= 2
	a *= 1
	a++
	a--

	fmt.Println("Atribuição =>", a)

	fmt.Println("Igual", "Verdade" == "Verdade")
	fmt.Println("Diferente", 2 != 3)
	fmt.Println(">", 2 > 3)
	fmt.Println("<", 2 < 3)
	fmt.Println("<=", 2 <= 3)
	fmt.Println(">=", 2 >= 3)

	x := time.Now()
	y := time.Now()

	fmt.Println("Equals", x.Equal(y))

	type Name struct {
		name string
	}

	n1 := Name{"Adriana"}
	n2 := Name{"Adriana"}

	fmt.Println("Struct", n1 == n2)

	fmt.Println(trabalho(true, false))

	//não tem operador ternário
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func trabalho(trab1, trab2 bool) (bool, bool, bool) {

	passear := trab1 && trab2

	estudar := trab1

	ficarEmCasa := !trab2

	return passear, estudar, ficarEmCasa

}
