// LOAD BALANCER

package main

import (
	"fmt"
	"time"
)

// TODA VEZ QUE MANDAR UM DADO PRO CANAL VAI LER E IMPRIMIR
func worker(workerID int, data chan int) {
	for x := range data { // ler do canal
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // THREAD 1
	canal := make(chan int)

	qtdWorkers := 1000000

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, canal)
	}

	// go worker(1, canal) // Thread 2 Go Routine
	// go worker(2, canal) // Thread 3 Go Routine
	// go worker(3, canal) // Thread 4 Go Routine...

	for i := 0; i < 1000000000; i++ {
		// tamo enchendo o canal com i
		canal <- i
	}
}

// Exemplo: streaming da Globo onde você pausa um vídeo e esse timestamp é salvo para retornar depois. São milhares de requisições. Com Golang os caras fazem isso com 500mb de memória com essa concorrência.
