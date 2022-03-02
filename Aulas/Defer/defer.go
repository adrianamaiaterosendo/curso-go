package main

import "fmt"

func valorPremio(num int) int {
	defer fmt.Println("Fim")
	if num > 500 {
		return 500
	} else {
		return num
	}
}

func main() {
	fmt.Println(valorPremio(600))
}
