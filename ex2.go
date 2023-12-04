package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	if len(slice) > 2 {
		slice = append(slice[:2], slice[3:]...)
	} else {
		fmt.Println("O slice não tem elementos suficientes para remover o terceiro.")
		return
	}

	fmt.Println("Slice resultante:", slice)
}
