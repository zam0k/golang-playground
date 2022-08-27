package main

import "fmt"

func main() {
	//declaração explicita
	var variavel1 string = "Variavel 1"

	//declaração implicita (inferencia de tipo)
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//declaração explicita de mais de uma variavel ao mesmo tempo

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	//declaração implicita de mais de uma variavel ao mesmo tempo

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, variavel6)

	// tudo isso abordado também vale para constante

	const constante = "Constante 1"
	fmt.Println(constante)

	//para inverter os valores não precisa de variavel auxiliar :o

	variavel5, variavel6 = "Variavel 6", "Variavel 5"
	fmt.Println(variavel5, variavel6)
}
