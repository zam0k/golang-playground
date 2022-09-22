package main

import (
	"fmt"
	"introducao-a-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")

	fmt.Println(tipoEndereco)
}