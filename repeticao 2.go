package main

import "fmt"

//Faça um algoritmo que imprima os números pares de 0 a 20.

func main() {
	i := 0

	for i < 21 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}

}
