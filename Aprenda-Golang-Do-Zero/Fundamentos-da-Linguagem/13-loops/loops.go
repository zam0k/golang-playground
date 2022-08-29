package main

import (
	"fmt"
	"time"
)

// Golang só tem o For como estrutura de repetição

func main() {
	i := 0

	for i < 3 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 3; j++ {
		fmt.Println("Incrementando", j)
	}

	//clausula range é para iterar

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	//em iteração de palavra ele traz o numero da tabela ascii
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Da Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// NÃO DA PRA FAZER MAP EM STRUCT

	// for {
	// 	fmt.Println("Executando infinitamente...")
	// }
}
