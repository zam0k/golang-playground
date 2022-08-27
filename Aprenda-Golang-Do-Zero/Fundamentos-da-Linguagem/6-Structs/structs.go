package main

import "fmt"

/*
	Não existe CLASSES em go, o mais proximo disso são structs
*/

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario

	u.nome = "Castelo"
	u.idade = 25

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	//forma mais rapida de inferir valores
	u2 := usuario{"Castelão", 25, enderecoExemplo}

	fmt.Println(u2)

	//para passar só um valor por algum motivo:
	u3 := usuario{nome: "Castelinho"}

	fmt.Println(u3)
}
