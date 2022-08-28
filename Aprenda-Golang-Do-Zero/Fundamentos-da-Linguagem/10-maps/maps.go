package main

import "fmt"

func main() {
	// # Maps
	// 				key		value
	usuario := map[int]string{
		1: "Fulano",
		2: "Da Silva",
	}

	fmt.Println(usuario)

	fmt.Println(usuario[1])

	/*
		map é bem mais rigido que structs, chave e valor tem que ser todos do mesmo tipo e o acesso é diferente
	*/

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Fulano",
			"ultimo":   "Da Silva",
		},
		"curso": {
			"campus": "Universidade da ",
			"nome":   "Engenharia",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
