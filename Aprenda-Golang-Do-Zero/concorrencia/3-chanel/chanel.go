package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo", canal)


	//ler mais sobre deadlocks
	// deadlock não é pego em compilação, só em execução
	// da pra verificar se um chanel está aberto ou fechado


	//canais são ações bloqueantes, ele deixa o programa parado até o canal receber um valor
	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem:= range canal {

		fmt.Println(mensagem)
	}

	fmt.Println("Fim")
}


func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	
	close(canal)
}