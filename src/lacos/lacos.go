package main

import "fmt"

func main() {
	nomes := []string{"Erick", "Luiz", "Maiko", "Marvin", "Nilson"}

	for i := 0; i < len(nomes); i++ {
		fmt.Println("Nomes no array", nomes[i])
	}

	fmt.Println("Outro laço\n.")

	for index, value := range nomes {
		fmt.Println(value, "Na posicao", index)
	}

	fmt.Println("Outro laço\n.")

	breaker := 0
	for {
		fmt.Println(nomes[breaker])

		if breaker >= len(nomes)-1 {
			break
		}

		breaker = breaker + 1
	}
}
