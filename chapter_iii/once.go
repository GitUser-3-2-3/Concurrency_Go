package main

import (
	"fmt"
	"sync"
)

func once1() {
	var count int

	increment := func() { count++ }
	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Println("count:", count)
}
