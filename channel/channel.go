package main

import (
	"fmt"
)

func main()  {
	closeClosed()
	sendToClosed()
	readFromClosed()
	selectRWBuf()
}

// closeClosed func show behavior what if try close a closed channel
func closeClosed()  {
	ch := make(chan int)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Try close already closed channel got panic: %v\n", r)
		}
	}()
	close(ch)
	close(ch)
}
// sendToClosed func show behavior what if try send on closed channel
func sendToClosed()  {
	ch := make(chan int)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Try write already closed channel got panic: %v\n", r)
		}
	}()
	close(ch)

	ch <- 42
}

// readFromClosed func show behavior what if try read on closed channel
func readFromClosed()  {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Try write already closed channel got panic: %v\n", r)
		}
	}()

	v, ok := <- ch
	fmt.Printf("Read from not closed channel: got value %v: okIdiom %t\n", v, ok)
	close(ch)
	for i := 0; i < 5; i++ {
		v, ok = <- ch
		fmt.Printf("From from closed channel in loop %d: got value %v: okIdiom %t\n", i, v, ok)
	}
}

func selectRWBuf()  {
	ch := make(chan int, 1) // non buf chan will block goroutine at send and make deadlock

	defer close(ch)

	for i := 0; i <= 10; i++ {
		select {
		case ch <- i:
			fmt.Printf("Send i %v to channel\n", i)
		case num := <-ch:
			fmt.Printf("Read i %v from channel\n", num)
		}
	}
}