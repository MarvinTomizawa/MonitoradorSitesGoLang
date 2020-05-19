package main

import (
	"fmt"
)

func main() {
	var valor int
	fmt.Println("Testando condições")
	fmt.Println("Informe um numero:")
	fmt.Scan(&valor)

	if valor == 1 {
		fmt.Println("O valor é um")
	} else {
		fmt.Println("O valor não é um")
	}

	switch valor {
	case 2:
		fmt.Println("Na verdade o valor é dois.")
	case 3:
		fmt.Println("Na verdade o valor é tres.")
	default:
		fmt.Println("O valor num é nem dois e nem tres")
	}
}
