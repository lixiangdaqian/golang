package channel3_test

import (
	"fmt"
	"sync"
	"testing"
)

//channel关闭
func DataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func DataProcesser(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if ret, ok := <-ch; ok {
				fmt.Println(ret)
			} else {
				fmt.Println("break")
				break
			}

		}
		wg.Done()
	}()
}

func TestChannel(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	DataProducer(ch, &wg)
	wg.Add(1)
	DataProcesser(ch, &wg)
	wg.Wait()
}
