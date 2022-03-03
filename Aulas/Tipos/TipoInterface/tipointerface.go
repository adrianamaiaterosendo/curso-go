package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	//consigo usar a variável coisa com vários tipos
	var coisa interface{}

	coisa = 3

	fmt.Println(coisa)

	coisa = "Curso Golang"

	fmt.Println(coisa)

	coisa = true

	fmt.Println(coisa)

	coisa = curso{
		nome: "Aprendendo Golang",
	}

	fmt.Println(coisa)

}
