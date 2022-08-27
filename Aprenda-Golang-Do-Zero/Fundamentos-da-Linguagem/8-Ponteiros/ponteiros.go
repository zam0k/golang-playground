package main

import "fmt"

func main() {
	var var1 int = 10
	// atribuição por cópia
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)

	/*
		Atribuição por ponteiros

		PONTEIROS são uma referência de memória

		*

	*/

	var var3 int = 100
	var ponteiro *int = &var3

	fmt.Println(var3, *ponteiro) //desreferenciação

	var3++
	fmt.Println(var3, *ponteiro)
}
