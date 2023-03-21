package main

//Faça um algoritmo que leia dois números inteiros e mostre o maior deles.

import "fmt"

func main() {
	var x int
	fmt.Println("Qual é o valor de x?")
	fmt.Scan(&x)

	var y int
	fmt.Println("Qual o valor de y?")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("x é maior que y")
	} else if x < y {
		fmt.Println("x é menor que y")
	} else {
		fmt.Println("x e y são iguais")
	}

}
