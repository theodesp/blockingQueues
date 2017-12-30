package blockingQueues

import (
	. "gopkg.in/check.v1"
	"sync"
	"testing"
)

type SimpleQueue interface {
	Put(item interface{}) (bool, error)
	Get() (interface{}, error)
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func benchmarkPut(c *C, writers int, readers int, q SimpleQueue) {
	wg := sync.WaitGroup{}

	for writer := 0; writer < writers; writer++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Put(i)
			}
			wg.Done()
		}(&wg)
	}

	for reader := 0; reader < readers; reader++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Get()
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}

func benchmarkPutMoreWriters(c *C, writers int, readers int, q SimpleQueue) {

	wg := sync.WaitGroup{}
	rest := writers - readers

	for writer := 0; writer < writers; writer++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Put(i)
			}
			wg.Done()
		}(&wg)
	}

	for reader := 0; reader < readers; reader++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Get()
			}
			wg.Done()
		}(&wg)
	}
	c.ResetTimer()
	for reader := 0; reader < rest; reader++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Get()
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()

}

func benchmarkPutMoreReaders(c *C, writers int, readers int, q SimpleQueue) {

	wg := sync.WaitGroup{}
	rest := readers - writers

	for writer := 0; writer < writers; writer++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Put(i)
			}
			wg.Done()
		}(&wg)
	}

	for reader := 0; reader < readers; reader++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Get()
			}
			wg.Done()
		}(&wg)
	}
	c.ResetTimer()

	for writer := 0; writer < rest; writer++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < c.N; i++ {
				q.Put(i)
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
