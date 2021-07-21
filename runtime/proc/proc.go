package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	var wg sync.WaitGroup

	fmt.Println("NUM CPU: ", runtime.NumCPU())
	fmt.Println("NUM GO: ",runtime.NumGoroutine())
	wg.Add(6)
	go longInc(1, &wg)
	go longInc(2, &wg)
	go longInc(3, &wg)
	go longInc(4, &wg)
	go longInc(5, &wg)
	go longInc(6, &wg)
	fmt.Println("NUM GO: ",runtime.NumGoroutine())

	wg.Wait()
}

func longInc(pos int8, wg *sync.WaitGroup)  {

	fmt.Println("start #", pos)
	var c int64
	for i := 0; i < 100000000; i++ {
		c++
	}
	fmt.Println("end #", pos)

	wg.Done()
}
