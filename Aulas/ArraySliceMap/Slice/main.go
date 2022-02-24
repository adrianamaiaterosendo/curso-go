package main

import "fmt"

func main() {
	//array
	a1 := [6]int{1, 2, 3, 4, 5, 6}

	//slice
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)

	s2 := a1[2:3]

	fmt.Println(s2)

	s2 = a1[:2]

	fmt.Println(s2)
}
