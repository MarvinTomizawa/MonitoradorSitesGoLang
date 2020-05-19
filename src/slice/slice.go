package main

import (
	"fmt"
	"reflect"
)

func main() {
	nomes := []string{"Marvin", "Massaru", "Tomizawa"}
	nomesArray := [3]string{"Marvin", "Massaru", "Tomizawa"}

	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println(nomes)
	fmt.Println("Tamanho:", len(nomes))
	fmt.Println("Capacidade:", cap(nomes))

	fmt.Println(".")

	fmt.Println(reflect.TypeOf(nomesArray))
	fmt.Println(nomesArray)
	fmt.Println("Tamanho:", len(nomesArray))
	fmt.Println("Capacidade:", cap(nomesArray))

	fmt.Println(".")

	nomes = append(nomes, "Teste")

	// Ao dar append em um slice a capacidade dele dobra!!!
	fmt.Println("Depois de dar append")
	fmt.Println(nomes)
	fmt.Println("Tamanho:", len(nomes))
	fmt.Println("Capacidade:", cap(nomes))
}
