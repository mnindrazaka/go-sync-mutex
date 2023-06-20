package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Value int
}

func (c *Counter) Add() {
	c.Value++
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.RWMutex
	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mtx.Lock()
				counter.Add()
				mtx.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value)
}
