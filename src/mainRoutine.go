package src

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	// Simula uma tarefa que demora 1 segundo
	time.Sleep(1 * time.Second)
}

func MainRoutine() {
	start := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Printf("Tempo total de execução: %v\n", time.Since(start))
}
