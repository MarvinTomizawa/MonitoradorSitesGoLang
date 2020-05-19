package main

import "fmt"

func main() {
	numeros := pegaNumeros()
	//numero1, numero2, numero3, numero4 := pegaNumeros() Não funciona
	fmt.Println(numeros)
	fmt.Println(numeros[1])
}

func pegaNumeros() [4]int {
	var numeros [4]int
	numeros[0] = 4
	numeros[1] = 9
	numeros[3] = 4
	return numeros
}
