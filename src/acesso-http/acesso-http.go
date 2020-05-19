package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Teste")
	site := "https://www.alura.com.br"

	response, erro := http.Get(site)

	fmt.Println("Response")
	fmt.Println(response)
	fmt.Println("Erro")
	fmt.Println(erro)

	fmt.Println("Status:", response.Status)
	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Body", response.Body)
	fmt.Println("Request", response.Request)
	fmt.Println("Header", response.Header)
	fmt.Println("Cookies", response.Cookies)
}
