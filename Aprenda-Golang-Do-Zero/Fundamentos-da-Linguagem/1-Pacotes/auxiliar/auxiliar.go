package auxiliar

import "fmt"

/*
	In Golang, there are no Public and Private keywords. But you can distinguish Public and Private by using the uppercase and lowercase of the name. The Public function can be called by other packages, but the Private function can only be called by the package that defines it.
*/

/*
	É boa pratica no go ter um comentário acima da função
	exportada explicando o que ela faz
*/
func Escrever() {
	fmt.Println("Escrevendo do arquivo auxiliar...")
	escrever2()
}
