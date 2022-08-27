package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    uint8
}

type estudante struct {
	//pra passar todos os campos da struct "pai", basta só por o tipo da struct aqui
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("\"Herança\"")

	p1 := pessoa{"João", "Pedro", 20, 178}

	fmt.Println(p1)

	e1 := estudante{p1, "Pedagogia", "Univesidade Federal do -"}

	fmt.Println(e1)

	fmt.Println(e1.nome)
}
