// uma funcao recursiva é uma função que chama a si propria
package main

import "fmt"

//precisa de uma condição de parada para evitar stack overflow
func fibonnacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}

// não é legal fazer função recursiva pra casos muito grandes

func main() {
	// sequencia de fibonacci

	// 1 1 2 3 5 8 13

	posicao := uint(5)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonnacci(i))
	}

}
