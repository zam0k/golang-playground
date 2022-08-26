package main

import (
	"fmt"
	"modulo/Aprenda-Golang-Do-Zero/Fundamentos-da-Linguagem/1-Pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("algo@gmail.com")
	fmt.Println(erro)
}
