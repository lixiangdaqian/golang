package ch26

import (
	"testing"
	"sync"
)

func TestPool(t *testing.T)  {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("create new object")
			return 1
		},
	}
	a := pool.Get().(int)
	t.Log(a)
	//不放回去，会重新创建一个对象
	//pool.Put(10)
	b := pool.Get().(int)
	t.Log(b)
}
