package main

import "fmt"

//recebe n ints, parece o spread operator eu acho
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

//da pra combinar

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

/*não da pra ter mais de um parametro variatico por função e
ele obrigatoriamente tem que ser o ultimo
*/

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6)
}
