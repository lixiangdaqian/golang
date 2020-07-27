package ch25

import (
	"fmt"
	"testing"
	"runtime"
	"time"
)

func all_done() int{
	num := 10
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			ch <- i
			fmt.Printf("all_done %d", i)
			fmt.Println()
		}(i)
	}
	finalRes := 0
	for i := 0; i < num; i++ {
		finalRes += <- ch
	}
	return finalRes

}

func TestAllDone(t *testing.T)  {
	t.Log("before gorutine num is ",runtime.NumGoroutine())
	finalRes := all_done()
	t.Log("result allDone is ", finalRes)
	time.Sleep(time.Second)
	t.Log("after gorutine num is ",runtime.NumGoroutine())
}