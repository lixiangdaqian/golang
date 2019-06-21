package ch7

import "testing"

func TestSliceInit(t *testing.T)  {
	var slice []int
	t.Log(len(slice), cap(slice))
	slice = append(slice, 1)
	t.Log(len(slice), cap(slice))
	slice1 := []int{1,2,3,4,5,6}
	t.Log(len(slice1), cap(slice1))
}

func TestSliceDefine(t *testing.T)  {
	slice := make([]int, 6, 10)
	slice = append(slice,1,2,3,4,5,6,7)
	t.Log(slice[9])
	t.Logf("len + %d + cap %d",len(slice), cap(slice))
	slice1 := slice[6:10]
	t.Logf("len + %d + cap %d",len(slice1), cap(slice1))
}