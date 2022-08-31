package main

import "fmt"

// aceita literalmente tudo independente de tipo
func generica(interf interface{}) {
	fmt.Println(interf)
}

/*
	melhor ter cautela na hora de usar esse tipo de ferramenta,
	afinal a tipagem forte ajuda em muita coisa
*/

func main() {
	generica("String")
	generica(1)
	generica(true)
}
