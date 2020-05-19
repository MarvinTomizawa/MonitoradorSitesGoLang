package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var entrada int
	for {
		mostraOpcoes()
		fmt.Scan(&entrada)

		switch entrada {
		case 1:
			monitorarSites()
		case 2:
			mostrarLogs()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
		}
	}
}

func mostraOpcoes() {
	fmt.Println("1- Monitorar sites")
	fmt.Println("2- Mostrar logs")
	fmt.Println("3- Sair do sistema")
}

func monitorarSites() {
	arquivo, err := os.Open("projeto/sites.txt")

	var sites []string

	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo de sites")
		return
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, _, err := leitor.ReadLine()
		if err == io.EOF {
			break
		}

		sites = append(sites, strings.TrimSpace(string(linha)))
	}

	for _, site := range sites {
		monitoraSite(site)
	}

	arquivo.Close()
}

func monitoraSite(site string) {
	logs, err := os.OpenFile("projeto/logs.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Houve um problema para salvar os logs")
	}

	status := getStatus(site)

	logs.WriteString(time.Now().Format("02/01/2006 15:04:05") + " Site:" + site + " status: " + strconv.FormatBool(status) + "\n")

	logs.Close()
}

func getStatus(site string) bool {
	response, err := http.Get(site)
	if err != nil {
		return false
	}

	if response.StatusCode == 200 {
		return true
	}

	return false
}

func mostrarLogs() {
	logs, err := ioutil.ReadFile("projeto/logs.txt")

	if err != nil {
		fmt.Println("Houve um problema para salvar os logs")
	}

	fmt.Println(string(logs))
}
