package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Olá, mundo!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("finotti.dev@gmail.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("finotti.dev")
	fmt.Println(erro2)
}
