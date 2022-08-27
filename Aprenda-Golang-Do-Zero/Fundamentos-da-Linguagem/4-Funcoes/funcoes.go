package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

/*
	As funções podem ter mais de um retorno (!!!!!!!!)
*/

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	/*
		Funções no golang também são tipos
	*/

	/*
		const c: string = (txt: string) => console.log(txt)
		c("algo aqui")
	*/

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSub := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSub)

	//_ permite que vc ignore um retorno

	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)

}
