package main

import (
	"fmt"
	"os"
	"strings"
)

const Len = 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println(" input err")
		return
	}
	for _, str :=range os.Args[1:]  {
		DiGui(str)
	}
	/*str1, str2 := os.Args[1], os.Args[2]
	DiGui(str1)
	DiGui(str2)*/
	/*res1 := process(str1)
	res2 := process(str2)
	res := append(res1, res2...)
	fmt.Println(res)*/
}

func DiGui(str string)  {
	if len(str) <= Len {
		fmt.Println(str + strings.Repeat("0", Len - len(str)))
		return
	}
	fmt.Println(str[:8])
	DiGui(str[8:])
}

func process(str string) []string {
	var result []string
	lTmp := 0
	sTmp := ""
	for _, s := range str {
		sTmp += string(s)
		lTmp += 1
		if lTmp == Len {
			result = append(result, sTmp)
			sTmp = ""
			lTmp = 0
			continue
		}
	}
	if lTmp != 0 {
		result = append(result, sTmp+strings.Repeat("0", Len-lTmp))
	}
	return result
}
