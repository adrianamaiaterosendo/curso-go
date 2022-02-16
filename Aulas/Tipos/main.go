package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Para saber o tipo de uma variável", reflect.TypeOf(38))

	var numerosPositivos uint64 = 1

	fmt.Println(numerosPositivos)

	x := 45.9

	fmt.Println(x, reflect.TypeOf(x))

	y := true

	fmt.Println(y, reflect.TypeOf(y))

	s := "Vamos ver o meu tamanho"

	fmt.Println(s, len(s))

	//tipos zeros

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

}
