package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup

	fmt.Println("NUM CPU: ", runtime.NumCPU())
	fmt.Println("NUM GO: ",runtime.NumGoroutine())
	wg.Add(10)
	go longInc(1, &wg)
	go longInc(2, &wg)
	go longInc(3, &wg)
	go longInc(4, &wg)
	go longInc(5, &wg)
	go longInc(6, &wg)
	go longInc(7, &wg)
	go longInc(8, &wg)
	go longInc(9, &wg)
	go longInc(10, &wg)
	fmt.Println("NUM GO: ",runtime.NumGoroutine())

	wg.Wait()
}

// longInc just func for doing some stuff which not allow to switch context in loop go < 1.14  (see https://golang.org/doc/go1.14#runtime)
//
// Go1.14 - Goroutines are now asynchronously preemptible.
// As a result, loops without function calls no longer potentially deadlock the scheduler or significantly delay garbage collection
func longInc(pos int8, wg *sync.WaitGroup)  {

	start := time.Now()

	var c int64
	for i := 0; i < 10000000000; i++ {
		c++
	}
	fmt.Printf("#%v time start %v - time end %v\n", pos, start, time.Since(start))

	wg.Done()
}
