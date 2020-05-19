package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var entrada string
	fmt.Println("Informe um valor para adicionar no arquivo")
	fmt.Scan(&entrada)

	adicionarLinha(entrada)

	arquivos := lerTodasLinha()

	for _, arquivo := range arquivos {
		fmt.Println(arquivo)
	}
}

func lerTodasLinha() []string {
	arquivo, err := os.Open("arquivos/nomes.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(-1)
	}

	leitor := bufio.NewReader(arquivo)
	var linhas []string

	for {
		linha, _, err2 := leitor.ReadLine()

		if err2 == io.EOF {
			break
		}

		linhas = append(linhas, strings.TrimSpace(string(linha)))
	}

	arquivo.Close()

	return linhas
}

func lerArquivoInteiro() {
	//arquivo, err := os.Open("arquivos/nomes.txt")
	arquivo, err := ioutil.ReadFile("arquivos/nomes.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		os.Exit(-1)
	}

	fmt.Println(string(arquivo)) // Converte array de bytes para string
}

func adicionarLinha(linha string) {
	arquivo, err := os.OpenFile("arquivos/nomes.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Teve um problema ao escrever no arquivo")
	}

	fmt.Println("Arquivo a ser inserido:", linha)

	verdadeiro := true

	arquivo.WriteString(linha + " " + strconv.FormatBool(verdadeiro) + "\n")

	arquivo.Close()
}
