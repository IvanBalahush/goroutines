package workerPool

import (
	"log"
	"sync"
	models "taska/pkg/parser"
)

type WorkerPool struct {
	Count  int
	Sender chan models.Restaurant
	Ender  chan bool
}

func NewWorkerPool(count int) *WorkerPool {
	return &WorkerPool{
		Count:  count,
		Sender: make(chan models.Restaurant, count*2),
		Ender:  make(chan bool),
	}
}

func (p *WorkerPool) Run(wg *sync.WaitGroup, handler func(author models.Restaurant)) {
	defer wg.Done()
	var shop models.Restaurant
	for {
		select {
		case shop = <-p.Sender:
			handler(shop)
		case <-p.Ender:
			//fmt.Println(<- p.Sender)
			log.Println("I am finish")
			return
		}
	}
}

func (p *WorkerPool) Stop() {
	for i := 0; i < p.Count; i++ {
		p.Ender <- false
	}
	close(p.Sender)
	close(p.Ender)
}
