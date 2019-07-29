package channel2_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case a := <-quit:
			fmt.Println(a, "quit")
			return
		}
	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 5)
}

func worker(done chan bool) {
	done <- true
	time.Sleep(time.Second)
}

func TestWorker(t *testing.T)  {
	done := make(chan bool, 1)
	worker(done)
	<-done
}
