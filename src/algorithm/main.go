package main

import (
	"flag"
	"fmt"
	"strconv"
)

var todoString = flag.String("s", "", "待处理的字符串")

// https://www.cnblogs.com/crazybuddy/p/5396554.html
func main() {
	flag.Parse()
	resString := ""
	resLen := 0
	tmpString := fmt.Sprintf("%s", *todoString)
	if tmpString != "" {
		continueS := ""
		continueL := 0
		for i, s := range tmpString {
			_, err := strconv.Atoi(string(s))
			if err != nil {
				//赋值并重置
				if continueL > 1 {
					resString += continueS
					if continueL > resLen {
						resLen = continueL
					}
				}
				continueS = ""
				continueL = 0
				continue
			}
			//是数字就+1
			continueS += string(s)
			continueL += 1
			//边界条件，如果已经是最后一位了，也+上
			if i == len(tmpString) - 1 && continueL > 1{
				resString += continueS
				if continueL > resLen {
					resLen = continueL
				}
			}
		}
	}
	formatPrint(resString, resLen)
}

func formatPrint(resString string, resLen int) {
	fmt.Println("resString is： " + resString + " continue len is ：" + strconv.Itoa(resLen))
}
