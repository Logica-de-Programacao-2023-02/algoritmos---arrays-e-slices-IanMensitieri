package main

import "fmt"

func main() {
	array := [10]float64{2.5, 6.7, 3.2, 8.1, 4.9, 7.3, 1.5, 9.8, 5.4, 10.0}

	var elementosMaioresQueCinco []float64

	for _, elemento := range array {
		if elemento > 5 {
			elementosMaioresQueCinco = append(elementosMaioresQueCinco, elemento)
		}
	}

	fmt.Println("Novo Slice com elementos maiores que 5:", elementosMaioresQueCinco)
}
