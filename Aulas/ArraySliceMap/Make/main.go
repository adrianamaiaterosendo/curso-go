package main

import "fmt"

func main() {
	s := make([]int, 10)

	s[9] = 12

	fmt.Println(s)

	s = make([]int, 10, 20)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(s)

	s2 := make([]int, 10, 20)

	s2Copy := make([]int, len(s2))
	copy(s2Copy, s2)

	fmt.Println(s2)

}
