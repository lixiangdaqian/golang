package interface2_test

import (
	"fmt"
	"testing"
)

func doSth(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Print("inter ", v)
		fmt.Println()
		break
	case string:
		fmt.Print("string ", v)
		fmt.Println()
		break
	default:
		fmt.Print("unknown type")
		fmt.Println()
		break
	}

	/*if i, ok := p.(int); ok {
		fmt.Print("inter ", i)
		fmt.Println()
		return
	}
	if str, ok := p.(string); ok {
		fmt.Print("string ", str)
		fmt.Println()
		return
	}
	fmt.Print("unknown type")
	fmt.Println()*/
}
func TestNullInter(t *testing.T) {
	doSth(10)
	doSth("hello啊")
	map1 := make([]int, 6, 10)
	doSth(map1)
}
