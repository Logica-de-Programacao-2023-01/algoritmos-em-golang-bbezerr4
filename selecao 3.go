package main

import "fmt"

//Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.

func main() {
	var x int
	fmt.Println("Qual o valor de x?")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("x é par")
	} else {
		fmt.Println("x é ímpar")
	}

}
