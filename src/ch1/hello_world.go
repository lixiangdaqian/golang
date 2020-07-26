package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	sql := "select pages_name,SUM(exp_num) as exp,Sum(click_num) as click,SUM(exp_num+click_num) as total from t_sptapp_content_exp_click where  content_id = ? group by pages_name"
	sql = strings.TrimSpace(sql)
	fmt.Println(sql)
}
