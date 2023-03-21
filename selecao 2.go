package main

import "fmt"

//Faça um algoritmo que leia três números inteiros e mostre o menor deles.

func main() {
	var x int
	fmt.Println("qual o valor de x?")
	fmt.Scan(&x)

	var y int
	fmt.Println("Qual o valor de y?")
	fmt.Scan(&y)

	var z int
	fmt.Println("Qual o valor de z?")
	fmt.Scan(&z)

	if x > y && x > z {
		fmt.Println("x é maior que y e z")
	} else if y > x && y > z {
		fmt.Println("y é maior que x e z")
	} else if z > x && z > y {
		fmt.Println("z é maior que x e y")
	} else {
		fmt.Println("x, y e z são iguais")
	}
}
