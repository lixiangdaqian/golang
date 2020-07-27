package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChuanXing(t *testing.T) {
	Task()
	OtherTask()
}

func Task() string {
	//fmt.Println("task")
	time.Sleep(2 * time.Second)
	return "taskString"
}

func OtherTask() {
	fmt.Println("otherTask")
	time.Sleep(time.Second)
}

func ChannelTask() chan string {
	str_chan := make(chan string)
	go func() {
		fmt.Println("start put in chan")
		ret := Task()
		str_chan <- ret
		fmt.Println("exit put the chan")
	}()
	return str_chan
}

func TestChannel(t *testing.T) {
	chanes := ChannelTask()
	OtherTask()
	fmt.Println(<-chanes)
	time.Sleep(time.Second)
}
