package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	media(6.5)
	media(2)

	if i := aleatorio(); i > 8 {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")

	}

	j := 1

	for j <= 10 {
		fmt.Println("Contando", j)
		j++
	}

	for i := 1; i < j; i++ {
		fmt.Println("Contando", i)
	}

	fmt.Println(notaParaDiploma(8))

	fmt.Println(hora())
}

func media(nota float64) {
	if nota > 6 {
		fmt.Println("Aprovado")

	} else if nota == 6 {
		fmt.Println("Recuparação")

	} else {
		fmt.Println("Reprovado")
	}

}

func aleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)

}

func notaParaDiploma(n int) string {
	nota := n

	switch nota {
	case 5:
		return "Recuperação"
	case 10:
		return "Aprovado com sucesso!"
	default:
		return "Nota inválida"

	}

}

func hora() string {
	hora := time.Now()

	switch {
	case hora.Hour() > 12:
		return "Boa Tarde!"
	case hora.Hour() < 12:
		return "Bom dia!"
	default:
		return "Que dia é hoje?"

	}

}
