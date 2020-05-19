package main

import "fmt"

func main() {
	mostraOpcoes()
	comando, valido := escolheComando()

	fmt.Println("Comando escolhido", comando)
	fmt.Println(valido)
}

func mostraOpcoes() {
	fmt.Println("1- Primeira opção")
	fmt.Println("2- Segunda opção")
	fmt.Println("3- terceira opção")
}

func escolheComando() (int, string) {
	comando := 0
	var valido string

	fmt.Scan(&comando)

	switch comando {
	case 1:
		valido = "Comando 1 valido"
	case 2:
		valido = "Comando 2 valido"
	default:
		valido = "Demais campos invalidos"
	}
	return comando, valido
}
