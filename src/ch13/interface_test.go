package ch13

import (
	"testing"
	"time"
	"fmt"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoPorgrammer struct {

}

func (goP *GoPorgrammer)WriteHelloWorld() string  {
	return "fmt.Println(\"hello world\")"
}

func TestTypeInter(t *testing.T)  {
	var p Programmer
	p = new(GoPorgrammer)
	t.Log(p.WriteHelloWorld())
}

//别名命名
type OpF func(op int) int
func TimeSpend(inner OpF) OpF {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend : ", time.Since(start).Seconds())
		return ret
	}
}

func TestTimeSpend(t *testing.T)  {
	ts := TimeSpend(JustReturn)
	ts(10)
}

func JustReturn(op int) int  {
	time.Sleep(time.Second * 1)
	return op
}