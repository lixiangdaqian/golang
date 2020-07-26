package main

import "fmt"

func main() {
	//''标识rune类型，int32类型，+标识数字相加
	fmt.Print('a' + 'b' + 'c' + 'd')
	fmt.Printf("%#v\n", str("hello"))
}

type str string

func (s str) GoString() string {
	return string(s) + "!"
}
