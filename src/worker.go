package src

import (
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simula uma tarefa que demora 1 segundo
	println(id)
	time.Sleep(1 * time.Second)
}
