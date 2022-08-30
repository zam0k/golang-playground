package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func main() {
	defer funcao1()
	// adia a execução da função até o ultimo momento possivel
	funcao2()
}

//defer é muito bom de usar com banco de dados, tipo fechar a conexão com o banco
