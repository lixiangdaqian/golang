package once_test

import (
	"fmt"
	"testing"
)

type Singleton struct {
}

var singleton Singleton

func getSingleton() {

}

func Test1(t *testing.T) {
	if errFunc() == nil {

		fmt.Println("errFunc 对了哦，外部接收到也是nil")
	} else {

		fmt.Println("errFunc 错了咦，外部接收到不是nil哦")

	}

	if rightFunc() == nil {

		fmt.Println("rightFunc 对了哦，外部接收到也是nil")
	} else {

		fmt.Println("rightFunc 错了咦，外部接收到不是nil哦")

	}

}

//错误：将nil的people给空接口后接口就不为nil,因为interface中的value为nil但type不为nil
func errFunc() *People {
	var p *People

	return p
}

//正确处理返回nil给接口的方式：直接将nil赋给interface
func rightFunc() IPeople {
	var p *People

	return p
}

type IPeople interface {
	hello()
}
type People struct {
}

func (p *People) hello() {
	fmt.Println("github.com/meetbetter")
}

func errFunc1(in int) *People {
	if in == 0 {
		fmt.Println("importantFunc返回了一个nil")
		return nil
	} else {
		fmt.Println("importantFunc返回了一个非nil值")
		return &People{}
	}

}

func Test222(t *testing.T)  {

	in := 0

	i := errFunc1(in)

	if i == nil {

		fmt.Println("errFunc1 哈，外部接收到也是nil")
	} else {

		fmt.Println("errFunc1 咦，外部接收到不是nil哦")
		fmt.Printf("%v, %T\n", i, i)
	}
}
