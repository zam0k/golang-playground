package main

import "fmt"

func main() {
	// Estruturas de controle

	// if/elseif/else -> nada de muito anormal

	numero := 10

	if numero > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor ou igual a 10")
	}

	/*
		no golang as chaves são obrigatórias, mesmo com if com uma linha só
	*/

	// if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que 0")
	} else {
		fmt.Println("Número menor que 0")
	}

	/*
		essa variavel é limitada ao escopo do if/else
		fmt.Println(outroNumero) <- error
	*/

}
