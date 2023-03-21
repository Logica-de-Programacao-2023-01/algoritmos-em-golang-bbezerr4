package main

//Faça um algoritmo que imprima os números ímpares de 1 a 19.

import "fmt"

func main() {
	i := 0

	for i < 20 {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
	}

}
