package map_test

import "testing"

func TestInitMap(t *testing.T)  {
	map1 := map[string]int{"string1":1, "string2":2}
	t.Log(map1["string1"])
	t.Logf("len:%d",len(map1))
	map2 := map[string]int{}
	map2["test1"] = 3
	t.Log(map2["test1"])
	t.Logf("len:%d",len(map2))
	map3 := make(map[string]int, 10)
	t.Logf("len:%d",len(map3))
}

func TestNotExistKeys(t *testing.T)  {
	map1 := map[int]int{}
	t.Log(map1[1])
	map1[2] = 0
	if _,ok := map1[2];ok {
		t.Log("ok")
	}else {
		t.Log("not ok")
	}
}
func TestMapRange(t *testing.T){
	map1 := map[string]int{"string1":1, "string2":2}
	for key,val := range map1 {
		t.Logf("keys is %s && value is %d", key, val)
	}
}
