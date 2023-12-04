package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var somaPares int

	for i, elemento := range array {
		if i%2 == 0 {
			somaPares += elemento
		}
	}

	fmt.Printf("Soma dos elementos nas posições pares: %d\n", somaPares)
}
