package goroutine_test

import (
	"testing"
	"time"
)

func TestGorutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//值传递，其实是复制了一份数据
		go func(i int) {
			t.Log(i)
		}(i)
		time.Sleep(300 * time.Microsecond)
	}
}
