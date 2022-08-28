package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string //todos sao obrigatoriamente do mesmo tipo
	fmt.Println(array1)
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}

	fmt.Println(array2)

	/*
		Array é bem inflexivel, o mais próximo de flexivel que conseguimos é:
	*/
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	/*
		Slice, que é flexivel e muito utilizado
	*/

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2) //cria uma fatia do array (in, ex)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)
}
