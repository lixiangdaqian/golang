package funAdv_test

import "testing"

//变长参数
func TestVarLongParam(t *testing.T) {
	r1 := Sum(1, 2, 3, 4)
	r2 := Sum(2, 3, 4, 5, 6, 7)
	t.Log(r1, r2)
}

func Sum(ops ...int) int {
	ret := 0
	for _, val := range ops {
		ret += val
	}
	return ret
}

//测试deffer
func TestDeffer(t *testing.T) {
	defer func() {
		t.Log("defer do sth")
	}()
	t.Log("start func TestDeffer")
	panic("SomeThing happend")
	//t.Log("after panic")
}
