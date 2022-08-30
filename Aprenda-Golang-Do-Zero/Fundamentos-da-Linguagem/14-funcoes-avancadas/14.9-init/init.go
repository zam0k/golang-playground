package main

import "fmt"

func init() {
	//funcao init pode ser colocada 1 vez por arquivo
	fmt.Println("funcao executada antes da func main")
}

func main() {
	// main é uma vez por programa
	fmt.Println("Eu sou a main")
}
