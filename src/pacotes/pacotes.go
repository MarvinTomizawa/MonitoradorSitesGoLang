package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	var variavel1 int
	fmt.Println("Valor da variavel:", reflect.TypeOf(variavel1))

	fmt.Println("Informe um valor valido para a variavel:")
	fmt.Scan(&variavel1)

	switch variavel1 {
	case 1:
		fmt.Println("Valor um é valido.")
	case 2:
		fmt.Println("Valor dois é valido.")
	case 3:
		fmt.Println("Valor tres é valido.")
	default:
		fmt.Println("Valor invalido.")
		os.Exit(-1)
	}

	fmt.Println("Verificando")

	for i := 0; i < 3; i++ {
		fmt.Print("*")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("Finalizado")

	os.Exit(0)
}
