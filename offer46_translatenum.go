package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，
//1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
//一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//示例 1:
//输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

func translateNum(num int) int {
	if num <= 9 {
		return 1
	}
	tmp_str := fmt.Sprintf("%d", num)
	var start, end, count int = 0, 0, 1
	for i := 0; i < len(tmp_str); i++ {
		start, end = end, count
		count=end
		if i == 0 {
			continue
		}
		dp := tmp_str[i-1 : i+1]
		if dp <= "25" && dp >= "10" {
			count += start
		}
	}
	return count
}

func main() {
	tmp_str := strconv.Itoa(12235)
	strconv.Atoi()
	for i := 0; i < len(tmp_str); i++ {
		fmt.Println(reflect.TypeOf(tmp_str[i]))
	}
	fmt.Println(reflect.TypeOf(tmp_str))
}
