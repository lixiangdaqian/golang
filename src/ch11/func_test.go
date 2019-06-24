package func_test

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func TestMultiRetFuc(t *testing.T) {
	first, two := returnMultivalue()
	t.Log(first, two)
}

func returnMultivalue() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10), rand.Intn(20)
}

func TestRunTime(t *testing.T) {
	tssP := timeSpend(Slow)
	tssP(10)
}

func Slow(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend : ", time.Since(start).Seconds())
		return ret
	}
}

func TestRandFunc(t *testing.T)  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		fmt.Println(x)
	}
}