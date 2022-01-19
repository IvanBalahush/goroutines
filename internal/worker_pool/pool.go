package workerPool

import (
	"sync"
	"taska/pkg/models"
)

type WorkerPool struct {
	WorkerCount int
	Jobs        chan models.Restaurant
	results     chan models.Restaurant
}

func NewPool(count int) *WorkerPool {
	return &WorkerPool{
		WorkerCount: count,
		Jobs:        make(chan models.Restaurant, count),
		results:     make(chan models.Restaurant, count),
	}
}

func (w *WorkerPool) Run(wg *sync.WaitGroup, handler func(restaurant models.Restaurant))  {
	defer wg.Done()
	var restaurant models.Restaurant
	for {
		select {
		case restaurant = <-w.Jobs:
			handler(restaurant)
		case <-w.results:
			w.Stop()
			return
		}
	}
}

func (w *WorkerPool) Stop()  {
	for i := 0; i < w.WorkerCount; i++ {
		w.results<-models.Restaurant{}
	}
}