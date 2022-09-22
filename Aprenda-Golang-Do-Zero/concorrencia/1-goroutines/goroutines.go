package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO
	go escrever("Olá Mundo") //goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
