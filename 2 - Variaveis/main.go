package main

import "fmt"

// Este programa demonstra diferentes formas de declarar e inicializar variáveis em Go.
//
// A função principal mostra:
// 1. Declaração explícita de variável com a palavra-chave `var` e especificação de tipo.
// 2. Declaração curta de variável usando o operador := (o tipo é inferido).
// 3. Declaração de múltiplas variáveis em um bloco var.
// 4. Declaração e inicialização de múltiplas variáveis em uma única linha.
// 5. Troca de valores de variáveis sem usar uma variável temporária.
//
// Cada método de declaração é seguido pela impressão das variáveis para demonstrar que foram
// inicializadas com sucesso.

func main() {
	var varaivel1 string = "Variável 1"
	fmt.Println(varaivel1)

	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
