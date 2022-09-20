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

func escrever(texto string) string {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
