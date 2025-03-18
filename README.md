# Introduction to Golang
Estudando Golang

## O que é Go?

Go (ou Golang) é uma linguagem de programação de código aberto desenvolvida pelo Google em 2007 e lançada em 2009. Foi criada por Robert Griesemer, Rob Pike e Ken Thompson com o objetivo de ser eficiente, com sintaxe simples e segura.

## Características principais

- **Compilada**: Gera código nativo de máquina para execução rápida
- **Tipagem estática**: Tipos são verificados em tempo de compilação
- **Garbage Collection**: Gerenciamento automático de memória
- **Concorrência nativa**: Goroutines e channels facilitam programação concorrente
- **Sintaxe simples**: Projetada para ser fácil de ler e escrever
- **Bibliotecas padrão robustas**: Muitas funcionalidades disponíveis nativamente

## Instalação e configuração

### Instalação

1. Acesse [golang.org/dl](https://golang.org/dl/)
2. Baixe a versão adequada para seu sistema operacional
3. Siga as instruções do instalador

### Verificação de instalação

```bash
go version
```

### Configurando o workspace

Por padrão, Go espera que seu código esteja em um diretório específico:

```
$HOME/go/         # Diretório base (GOPATH)
  |-- bin/        # Executáveis compilados
  |-- pkg/        # Pacotes compilados
  |-- src/        # Código fonte
```

## Comandos básicos do Go

- `go run arquivo.go`: Compila e executa o código
- `go build arquivo.go`: Compila e gera um executável
- `go install`: Compila e instala o pacote
- `go fmt`: Formata código automaticamente
- `go test`: Executa testes
- `go mod init nome_módulo`: Inicializa um novo módulo
- `go get pacote`: Download e instala dependências

## Sintaxe básica

### Olá Mundo

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá, Mundo!")
}
```

### Tipos de dados básicos

- **Números inteiros**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, etc.
- **Números de ponto flutuante**: `float32`, `float64`
- **Texto**: `string`
- **Booleanos**: `bool` (true/false)
- **Complexos**: `complex64`, `complex128`

### Variáveis

```go
// Declaração simples
var nome string = "Maria"

// Inferência de tipo
idade := 30

// Constantes
const PI = 3.14159
```

### Estruturas de controle

#### Condicionais

```go
if idade >= 18 {
    fmt.Println("Maior de idade")
} else {
    fmt.Println("Menor de idade")
}
```

#### Switch

```go
switch diaDaSemana {
case "sábado", "domingo":
    fmt.Println("Final de semana!")
default:
    fmt.Println("Dia útil.")
}
```

#### Loops

```go
// For tradicional
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// For como "while"
contador := 0
for contador < 10 {
    contador++
}

// For infinito
for {
    // precisa de um break para sair
}
```

### Estruturas de dados

#### Arrays

```go
var numeros [5]int // Array de 5 inteiros

// Inicializando
nomes := [3]string{"João", "Maria", "Pedro"}
```

#### Slices (mais comuns que arrays)

```go
numeros := []int{1, 2, 3, 4, 5}
numeros = append(numeros, 6) // Adiciona elemento
```

#### Maps

```go
pessoa := map[string]string{
    "nome":      "João",
    "sobrenome": "Silva",
}
```

### Funções

```go
// Função simples
func soma(a int, b int) int {
    return a + b
}

// Múltiplos retornos
func divisao(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

### Structs (tipos personalizados)

```go
type Pessoa struct {
    Nome     string
    Idade    int
    Endereco string
}

// Criando uma pessoa
joao := Pessoa{
    Nome:     "João Silva",
    Idade:    30,
    Endereco: "Rua ABC, 123",
}
```

## Concorrência em Go

### Goroutines

```go
// Função executada concorrentemente
go minhaFuncao()
```

### Channels (canais)

```go
// Criando um canal
canal := make(chan string)

// Enviando dados para o canal
go func() {
    canal <- "Olá do canal!"
}()

// Recebendo dados do canal
mensagem := <-canal
fmt.Println(mensagem)
```

## Tratamento de erros

```go
arquivo, erro := os.Open("arquivo.txt")
if erro != nil {
    fmt.Println("Erro ao abrir arquivo:", erro)
    return
}
defer arquivo.Close()
```

## Recursos para aprendizado

- [Tour of Go](https://tour.golang.org/)
- [Documentação oficial](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)

## Próximos passos

- Aprenda sobre pacotes e módulos
- Explore a biblioteca padrão
- Pratique com pequenos projetos
- Estude padrões de concorrência avançados

## Sugestões de Projetos Iniciais

Aqui estão algumas ideias de projetos para praticar seus conhecimentos em Go:

### 1. CLI Todo List
Desenvolva um aplicativo de linha de comando para gerenciar tarefas:
- Adicionar, remover e marcar tarefas como concluídas
- Armazenar tarefas em um arquivo JSON ou SQLite
- Implementar filtragem por status ou categoria
- Usar pacotes como `cobra` ou `urfave/cli` para interface CLI

### 2. API REST Simples
Crie uma API para um recurso simples (como produtos ou usuários):
- Implemente operações CRUD (Create, Read, Update, Delete)
- Use o pacote nativo `net/http` ou frameworks como `gin-gonic/gin` ou `gorilla/mux`
- Conecte a um banco de dados (MySQL, PostgreSQL ou MongoDB)
- Adicione validação de entrada e tratamento de erros

### 3. Crawler/Scraper Web
Desenvolva um programa para extrair informações de websites:
- Use o pacote `net/http` para fazer requisições HTTP
- Parse HTML com `goquery` (similar ao jQuery)
- Processe dados em paralelo usando goroutines
- Salve os resultados em arquivo ou banco de dados

### 4. Chat em Tempo Real
Crie um simples chat cliente-servidor:
- Use WebSockets para comunicação bidirecional
- Implemente salas de chat e mensagens privadas
- Armazene histórico de mensagens temporariamente
- Crie uma interface web simples para interação

### 5. Gerador de Imagens
Desenvolva uma ferramenta para manipulação de imagens:
- Redimensione imagens usando o pacote `image`
- Aplique filtros simples (como preto e branco, sépia)
- Adicione texto ou marca d'água nas imagens
- Crie interface CLI para executar as operações

### 6. Monitoramento de Sistema
Crie uma ferramenta para monitorar recursos do sistema:
- Colete métricas como uso de CPU, memória e disco
- Implemente alertas para métricas acima de limites definidos
- Apresente estatísticas em uma interface web simples
- Utilize goroutines para coleta paralela de dados

### 7. Serviço de Encurtador de URLs
Desenvolva um serviço para encurtar URLs:
- Gere códigos curtos para URLs longas
- Implemente redirecionamento
- Armazene dados em banco de dados
- Adicione estatísticas de acesso para cada URL

### 8. Biblioteca Simples
Desenvolva um sistema de gerenciamento de biblioteca:
- Cadastro de livros, usuários e empréstimos
- Interface CLI ou web simples
- Pesquisa e filtragem de livros
- Notificações de devolução

Cada projeto pode ser expandido à medida que você ganha confiança. Comece com versões simplificadas e vá adicionando recursos conforme se familiariza com a linguagem e suas bibliotecas.
