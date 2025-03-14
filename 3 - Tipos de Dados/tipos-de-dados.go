package main

import (
	"errors"
	"fmt"
	"unsafe"
)

// Este arquivo demonstra os tipos de dados fundamentais em Go, incluindo exemplos práticos
// para cada tipo. O código explora desde tipos básicos como booleanos e números inteiros
// até tipos compostos como slices, mapas e interfaces.
//
// Tipos abordados:
//
// Tipos básicos:
//   - Boolean: representação de valores lógicos verdadeiro/falso
//   - Numéricos:
//     * Inteiros com sinal: int, int8, int16, int32, int64
//     * Inteiros sem sinal: uint, uint8, uint16, uint32, uint64, uintptr
//     * Aliases: byte (uint8), rune (int32)
//     * Ponto flutuante: float32, float64
//     * Complexos: complex64, complex128
//   - String: sequências de caracteres UTF-8
//   - Error: tipo para representar condições de erro
//
// Tipos compostos:
//   - Array: coleção de tamanho fixo
//   - Slice: coleção de tamanho dinâmico
//   - Map: coleção de pares chave-valor
//   - Struct: agrupamento de campos com tipos potencialmente diferentes
//   - Ponteiro: referência a um valor na memória
//   - Interface: definição de comportamentos
//   - Canal: mecanismo para comunicação entre goroutines
//   - Função: tipo que representa um comportamento executável
//
// Conceitos especiais:
//   - nil: valor zero para ponteiros, interfaces, maps, slices, channels e funções
//   - Zero values: valores padrão para variáveis não inicializadas

func main() {
	// Tipo Booleano
	var boolExample bool = true
	fmt.Println("Booleano:", boolExample)

	// Tipos numéricos - Inteiros
	var intExample int = 42
	var int8Example int8 = 127                   // -128 até 127
	var int16Example int16 = 32767               // -32768 até 32767
	var int32Example int32 = 2147483647          // -2147483648 até 2147483647
	var int64Example int64 = 9223372036854775807 // -9223372036854775808 até 9223372036854775807

	// Inteiros sem sinal
	var uintExample uint = 42
	var uint8Example uint8 = 255                    // 0 até 255
	var uint16Example uint16 = 65535                // 0 até 65535
	var uint32Example uint32 = 4294967295           // 0 até 4294967295
	var uint64Example uint64 = 18446744073709551615 // 0 até 18446744073709551615

	// Tipos inteiros dependentes da máquina
	var uintptrExample uintptr = uintptr(unsafe.Pointer(&intExample))

	fmt.Println("Inteiros:", intExample, int8Example, int16Example, int32Example, int64Example)
	fmt.Println("Inteiros sem sinal:", uintExample, uint8Example, uint16Example, uint32Example, uint64Example)
	fmt.Println("Uintptr:", uintptrExample)

	// Byte e rune (apelidos)
	var byteExample byte = 'A' // apelido para uint8
	var runeExample rune = 'Ж' // apelido para int32, representa um ponto de código Unicode
	fmt.Printf("Byte: %v (ASCII: %c), Rune: %v (Unicode: %c)\n", byteExample, byteExample, runeExample, runeExample)

	// Ponto flutuante
	var float32Example float32 = 3.14
	var float64Example float64 = 3.141592653589793
	fmt.Println("Ponto flutuante:", float32Example, float64Example)

	// Números complexos
	var complex64Example complex64 = 1 + 2i
	var complex128Example complex128 = 3 + 4i
	fmt.Println("Complexos:", complex64Example, complex128Example)
	fmt.Println("Parte real:", real(complex128Example), "Parte imaginária:", imag(complex128Example))

	// String
	var stringExample string = "Olá, Go!"
	fmt.Println("String:", stringExample)
	fmt.Println("Comprimento da string:", len(stringExample))
	fmt.Printf("Primeiro caractere: %c\n", stringExample[0])

	// Tipos compostos
	// Array - tamanho fixo
	var arrayExample [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arrayExample)

	// Slice - tamanho dinâmico
	var sliceExample []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", sliceExample, "Comprimento:", len(sliceExample), "Capacidade:", cap(sliceExample))

	// Map - pares chave-valor
	var mapExample map[string]int = map[string]int{"um": 1, "dois": 2, "três": 3}
	fmt.Println("Map:", mapExample)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	var structExample Person = Person{Name: "Alice", Age: 30}
	fmt.Println("Struct:", structExample)

	// Ponteiro
	var pointerExample *int = &intExample
	fmt.Println("Ponteiro:", pointerExample, "Valor:", *pointerExample)

	// Interface
	var interfaceExample any = "Posso armazenar qualquer tipo"
	fmt.Println("Interface:", interfaceExample)
	interfaceExample = 42
	fmt.Println("Interface após mudança:", interfaceExample)

	// Canal - para comunicação entre goroutines
	channelExample := make(chan int)
	go func() { channelExample <- 10 }()
	fmt.Println("Valor do canal:", <-channelExample)

	// Tipo Error - para representação de erros
	var errExample error = errors.New("isto é um exemplo de erro")
	fmt.Println("Error:", errExample)

	// Criando um erro com mensagem formatada
	nome := "usuário"
	senhaInvalida := fmt.Errorf("senha inválida para %s", nome)
	fmt.Println("Erro formatado:", senhaInvalida)

	// Verificando se um erro ocorreu
	var errNulo error // valor zero é nil (equivalente a null)
	if errNulo == nil {
		fmt.Println("Sem erro: errNulo é nil")
	}

	// Valores nil (equivalente a null em outras linguagens)
	// Em Go, apenas tipos de referência podem ser nil:
	var ponteiro *int = nil             // ponteiro nil
	var fatia []int = nil               // slice nil
	var dicionario map[string]int = nil // map nil
	var canal chan int = nil            // canal nil
	var funcao func() = nil             // função nil
	var i any = nil                     // interface nil

	fmt.Println("Nil em Go:")
	fmt.Println("- Ponteiro nil:", ponteiro)
	fmt.Println("- Slice nil:", fatia)
	fmt.Println("- Map nil:", dicionario)
	fmt.Println("- Canal nil:", canal)
	fmt.Println("- Função nil:", funcao == nil) // Corrigido: verificar se é nil em vez de imprimir a função
	fmt.Println("- Interface nil:", i)

	// Valores zero (equivalente a "undefined" em outras linguagens)
	// Go não tem undefined, mas toda variável declarada é inicializada com seu "valor zero"
	var intNaoInic int       // valor zero: 0
	var floatNaoInic float64 // valor zero: 0.0
	var boolNaoInic bool     // valor zero: false
	var stringNaoInic string // valor zero: "" (string vazia)
	var structNaoInic Person // valor zero: todos os campos com seus respectivos valores zero

	fmt.Println("\nValores zero em Go (equivalente a undefined):")
	fmt.Println("- Int não inicializado:", intNaoInic)
	fmt.Println("- Float não inicializado:", floatNaoInic)
	fmt.Println("- Bool não inicializado:", boolNaoInic)
	fmt.Println("- String não inicializada:", stringNaoInic)
	fmt.Println("- Struct não inicializada:", structNaoInic)

	// Tipo função
	var functionExample func(int, int) int = func(a, b int) int { return a + b }
	fmt.Println("Resultado da função:", functionExample(5, 3))
}
