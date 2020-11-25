package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，
// 返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//输入: num1 = "2", num2 = "3"
//输出: "6"

//示例 2:
//输入: num1 = "123", num2 = "456"
//输出: "56088"

//说明：
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var index1 = len(num1)
	var index2 = len(num2)

	num01 := []byte(num1)
	num02 := []byte(num2)

	tmp := make([]int, index1+index2)

	for i := 0; i < index1; i++ {
		for j := 0; j < index2; j++ {
			tmp[i+j+1] += int(num01[i]-'0') * int(num02[j]-'0')
		}
	}

	fmt.Println("tmp1",tmp)
	//处理进位
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i]/10
		tmp[i] = tmp[i] % 10
	}
	fmt.Println("tmp2",reflect.TypeOf(tmp),tmp)
	//去掉首位为0
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}

	result := ""
	for i := 0; i < len(tmp); i++ {
		result += strconv.Itoa(tmp[i])
		fmt.Println(result,"sss",result[i])
	}

	return string(result)
}

func main() {
	num1, num2 := "498828660196", "840477629533"
	s := multiply(num1, num2)
	fmt.Println(reflect.TypeOf(s), s)
}
