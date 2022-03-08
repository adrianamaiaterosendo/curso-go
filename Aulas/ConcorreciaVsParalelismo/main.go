package main

import (
	"fmt"
	"time"
)

//Concorrência vs Papalelismo
//Paralelismo(Fazer duas coisas exatamente ao mesmo tempo) - vários processadores - executar codigos simultaneamente
//Concorrência é a capacidade de administrar várias tarefas em um processador
//Go é muito inteligente para usar concorrência
//Em certos casos usar a concorrência é muito mais inteligente que o paralelismo pois quando se divide as tarefas para
//vários processadores depois eu tenho que juntar e paralelismo exige mais do hardware

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Por que vc não fala comigo", 3)
	//fale("Joao", "Tenho que esperar vc falar", 1)

	go fale("Maria", "Ouuuuu", 300)
	go fale("Joao", "Eiiiii", 300)

	time.Sleep(time.Second * 5)
	fmt.Println("Fim!")
}
