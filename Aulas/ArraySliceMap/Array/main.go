package main

import "fmt"

func main() {
	total := 0.0
	var notas [3]float64
	notas[0], notas[1], notas[2] = 9, 8, 10

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Println(media)

	numeros()
}

func numeros() {
	numeros := [...]int{1, 2, 3, 4, 5}

	//pode se ignorar um dos elementos utilizando _
	for i, num := range numeros {
		fmt.Println("Numero e índice", i, num)
	}
}
