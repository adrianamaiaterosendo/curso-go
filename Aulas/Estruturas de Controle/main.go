package main

import "fmt"

func main() {

	media(6.5)
	media(2)

}

func media(nota float64) {
	if nota >= 6 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")

	}

}
