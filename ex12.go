package main

import "fmt"

func main() {
	array := [5]int{2, 6, 9, 12, 7}

	var multiplosDeTres []int

	for _, elemento := range array {
		if elemento%3 == 0 {
			multiplosDeTres = append(multiplosDeTres, elemento)
		}
	}

	fmt.Println("Novo Slice com múltiplos de 3:", multiplosDeTres)
}
