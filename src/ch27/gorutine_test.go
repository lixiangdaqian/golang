package ch27

import (
	"testing"
	"context"
	"time"
	"fmt"
)

//结束gorutine
/*func TestStopGorutine(t *testing.T) {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控结束了")
				return
			default:
				fmt.Println("监控持续中")
				time.Sleep(time.Second * 2)
			}
		}
	}()
	fmt.Println("先等待6秒钟")
	time.Sleep(time.Second * 6)
	stop <- true
}*/

/*func TestStopWithContext(t *testing.T)  {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控结束了")
				return
			default:
				fmt.Println("监控持续中")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 6)
	fmt.Println("先等待6秒钟")

	cancel()
	time.Sleep(time.Second * 2)
}*/
//监控多个协程
func TestMultiGorutine(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch("协程1", ctx)
	go watch("协程2", ctx)
	go watch("协程3", ctx)
	time.Sleep(time.Second * 6)
	println("可以结束了")
	cancel()
	fmt.Printf("%s", ctx.Err())
	fmt.Println()
	time.Sleep(time.Second * 2)
}

func watch(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			println(name + "监控结束")
			return
		default:
			println(name + "持续监控中")
			time.Sleep(time.Second * 2)
		}
	}
}
