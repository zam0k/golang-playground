package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		números inteiros:

		int8, int16, int32, int64, int

		a diferença entre eles é o número de bits que eles suportam, se atribuido um numero excedente, eles retornam um integer overflow

		o int sozinho usa a arquitetura do seu pc como base

	*/

	var numero int16 = 100
	fmt.Println(numero)

	/*
		uint - Unsigned Integer
		int8, int16, int32, int64, int
	*/

	var numero2 uint32 = 100
	fmt.Println(numero2)

	// alias

	// rune == int32
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte == uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	/*
		números reais:

		float32, float64

	*/

	var numeroReal1 float32 = 123.456
	var numeroReal2 float64 = 1230000000.456

	fmt.Println(numeroReal1, numeroReal2)

	numeroReal3 := 123456.66

	fmt.Println(numeroReal3)

	/*
		Strings

		Não tem CHAR em go, eles são convertidos pra numero???
	*/

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println("numero do character na tabela ascii:", char)

	/*
		Valor 0 - Valor atribuido a uma variavel se você não atribuir a ela um valor

		string -> string vazia
		numeros -> 0
		error -> <nil>
		boolean -> false

	*/

	var texto string
	fmt.Println(texto)

	/*
		Boolean

		true or false
	*/

	var boolean bool
	fmt.Println(boolean)

	/*
		error
	*/

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)

}
