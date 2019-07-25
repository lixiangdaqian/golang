package sync_lock_test

import (
	"testing"
	"sync"
	"time"
)

//没有锁，共享内存，会产生竞争
func TestLock(t *testing.T) {
	counter := 0
	for i := 0; i < 5000;i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log("TestLock", counter)
}

//加锁，线程安全的
func TestLockThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000;i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)//为了等待所有的协程执行完
	t.Log("TestLockThreadSafe", counter)
}

//线程安全的
func TestLockThreadSafeWG(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000;i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()//会阻塞，知道wait的所有任务都执行完成
	//time.Sleep(1 * time.Second)
	t.Log("TestLockThreadSafe", counter)
}
