# GoLang

## Estrutura de pastas
> bin (Compilado)
> pkg (Pacotes utilizados)
> src (Código fonte do projeto)

## Comandos 
Buildar um compilavel GO e gera um executavel
```shell
go build hello.go
```

Buildar e executar 
```shell
go run hello.go
```

## Curiosidades
- Chave sempre do lado da função
```go
func go() {
    //Código
}
```

- Ponto e virgula opcional :c 

- Variaveis declaradas, primeiro nome, depois o tipo
```go
var nome string = "Marvin"
```
- Não deixa executar se existe alguma variável sem utilizaroooi

- Atribuições curtas com `:=`
```go
var variavel1 int = 10
variavel2 := 20
```
- if não utiliza parenteses ;^; 
```go
condicao := true
if condicao {
    fmt.Println("Que absurdo")
}
```
- Funções podem ter mais de um retorno
```go
func retornaIntEString() (int, string) {
    return 1, "Teste"
}

func main(){
    inteiro, texto := retornaIntEString()

    //Ignora o primeiro retorno no operador de identificador em branco
    _, texto2:= retornaIntEString()
}
```

- Não existe while, utiliza-se o for com argumentos diferentes
```go
for {
    fmt.Println("Escrevendo para sempre")
}
```

- Slices são abstrações de array para poder manipular sem se preocupar com tamanhos do array
    - Algumas funções que auxiliam a utilização de array são : `append, len, cap`
    - Ao adicionar um item com append em um slice a capacidade dele dobra caso a capacidade seja menor que a quantidade de items
```go
    slice := []int {"Teste", "Slice"}
    array := [3]int 

    fmt.Println("Tamanho:", len(slice)) //Retorna o tamanho do array: 3
    fmt.Println("Capacidade:", cap(slice)) // Retorna a capacidade do array: 3
    
    slice = append(slice, "Teste") // slice = {"Teste", "Slice", "Teste"}

    fmt.Println("Tamanho:", len(slice)) //Retorna o tamanho do array: 4
    fmt.Println("Capacidade:", cap(slice)) // Retorna a capacidade do array: 6
```
## Pacotes
fmt
- Entrada e saida de dados

os
- Conversa com o sistema operacional

reflect
- Consegue tratar um pouco reflection do sistema

time 
- Controle sobre datas e horas