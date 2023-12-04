package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)

	var numero int
	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&numero)

	estaPresente := false
	for _, elemento := range slice {
		if elemento == numero {
			estaPresente = true
			break
		}
	}

	if !estaPresente {
		slice = append(slice, numero)
	}

	fmt.Println("Slice resultante:", slice)
}
