package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("err input")
		return
	}
	strInputs := os.Args[1:]
	var inputs []int
	for _, input := range strInputs {
		intIp, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("err input: " + input)
			return
		}
		inputs = append(inputs, intIp)
	}
	Combine(inputs)
}

func Combine(inputs []int) {
	var res []string
	for _, v := range inputs {
		res = Padding(NumMapping[v], res)
	}
	fmt.Println(res)
}

func Padding(nums []string, res []string) []string {
	if len(res) == 0 {
		res = append(res, "")
	}
	var ret []string
	for _, r := range res {
		for _, n := range nums {
			ret = append(ret, r+n)
		}
	}
	return ret
}

var NumMapping = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "u"},
	{"w", "x", "y", "z"},
}
