package main

import "fmt"

//Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.

func main() {
	i := 0

	for i < 31 {
		if i%3 == 0 {
			fmt.Println(i)

		}
		i++
	}
}
