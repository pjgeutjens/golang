package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	WaitGroups()
	Channels()
	Selects()
	Looping()
}

func WaitGroups() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()

	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()
	fmt.Println("This is synchronous")
	wg.Wait()
}

func Channels() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go func() {
		ch <- "a message in a channel"
	}()
	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

func Selects() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	// pause execution for 10 milliseconds so the scheduler can see both goroutines
	time.Sleep(10 * time.Millisecond)

	// select statements will evaluate each case and execute a random matching case
	// select statements are blocking if they dont have a default case
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default: // this is a non-blocking select statement
		fmt.Println("no message received")
		// to active this case remove the goroutines above
	}
}

func Looping() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // without this the range will block the goroutine
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
