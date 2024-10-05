package main

// document.getElementsByClassName("product-item")
import (
	"math/rand"
	"sync"
	"time"
)

type ThreadControl struct {
	channel       chan int
	wg            *sync.WaitGroup
	mu            *sync.Mutex
	activeThreads int
}

func NewThreadControl(maxThreads int) ThreadControl {
	ret := ThreadControl{
		wg:            &sync.WaitGroup{},
		mu:            &sync.Mutex{},
		activeThreads: 0,
	}
	if maxThreads > 0 {
		ret.channel = make(chan int, maxThreads)
	}
	return ret
}

func (this *ThreadControl) Open() {
	this.mu.Unlock()
}

func (this *ThreadControl) Lock() {
	this.mu.Lock()
}

func (this *ThreadControl) Add(delta int) {
	this.channel <- 0
	this.activeThreads += delta
	this.wg.Add(delta)
}

func (this *ThreadControl) Done() {
	this.activeThreads--
	this.wg.Done()
	<-this.channel
}

var (
	count int
	mu    sync.Mutex // Mutex para proteger o acesso à variável count
)

func worker(wg *sync.WaitGroup) {
	mu.Lock()
	count++
	mu.Unlock()

	time.Sleep(time.Duration(rand.Float64() * 100000000000)) // Simula algum trabalho

	mu.Lock()
	count--
	mu.Unlock()
	wg.Done()
}

func main() {
	const maxGoroutines = 3000
	var wg sync.WaitGroup
	sem := make(chan int, maxGoroutines)

	go func() {
		for {
			println("Active goroutines:", count)
			time.Sleep(3 * time.Second)
		}
	}()

	for i := 0; ; i++ {
		wg.Add(1)
		sem <- 0

		go func() {
			worker(&wg)
			<-sem
		}()
	}

	wg.Wait() // Isso nunca será atingido, pois o loop é infinito
}
