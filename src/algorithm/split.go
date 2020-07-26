package main

import (
	"fmt"
	"os"
	"regexp"
)

func Split()  {
	str := os.Args[1]
	//mp := strings.Split(str, `\s+`)
	reg,_ := regexp.Compile(`\s+`)
	mp := reg.Split(str, -1)
	fmt.Println(mp,len(mp))
}

func main()  {
	Split()
}
