package ch25

import (
	"fmt"
	"testing"
	"runtime"
	"time"
)

//任意函数返回一个就认为是成功了,这种方式存在协程泄露，导致系统资源耗尽
func testAny() int{
	num := 10
	ch := make(chan int, num)
	for i := 1; i <= num; i++ {
		go func(i int) {
			ch <- i
			fmt.Printf("test any %d", i)
		}(i)
	}
	return <-ch
}

func TestAy(t *testing.T)  {
	t.Log("before:",runtime.NumGoroutine())
	testAny()
	time.Sleep(time.Second)
	t.Log("after:",runtime.NumGoroutine())
}
