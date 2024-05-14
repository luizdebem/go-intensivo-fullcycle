package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() { // Thread 1
	canal := make(chan string)

	go func() { // Thread 2
		canal <- "opa"
	}()

	msg := <-canal
	fmt.Println(msg) // Thread 1. Thread 2 se comunicando com a thread 1 pelo canal
}

// Processo -> alocar um bloco de memória
// Thread -> acessar o bloco de memória
// T1 e T2 -> acessam o mesmo bloco de memória
// race condition -> condição de corrida

// Go Routine -> Canal -> Go Routine 2
// INPUT -> OUTPUT
