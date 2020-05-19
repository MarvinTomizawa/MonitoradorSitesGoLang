// Import do package main
package main

import "fmt"

//declara-se funções com func
func main() {
	nomePessoa := "Marvin"
	var versao float32 = 1.1
	fmt.Println("Ola mundo")
	fmt.Println("Nome:", nomePessoa)
	fmt.Println("Versão aplicativo:", versao)

	fmt.Println("1- Iniciar Monitoramento.")
	fmt.Println("2- Exibir logs.")
	fmt.Println("0- Sair do Programa.")
	fmt.Println(".")

	var comando int
	fmt.Scan(&comando)
	// ou fmt.Scanf("%d", &comando)

	fmt.Println("Comando escolhido:", comando)
}
