package arr_test

import "testing"

func TestArr(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{11, 22, 33, 44, 55, 66}
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1)
	t.Log(arr2)

}

func TestArrRange(t *testing.T) {
	arr2 := [...]string{"1", "2", "3", "4", "5"}
	for str, _ := range arr2[1:3] {
		t.Log(str)
	}
}
func TestArrSec(t *testing.T) {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1_sec := arr1[3:len(arr1)]
	t.Log(arr1_sec)
}
