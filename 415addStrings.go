package main

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
//注意：
//
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。
//你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

func addStrings(num1 string, num2 string) string {

	if num1 == "0" && num2 == "0" {
		return "0"
	}
	var index1 = len(num1) - 1
	var index2 = len(num2) - 1
	var left int
	var result string
	for index1 >= 0 && index2 >= 0 {
		num01 := num1[index1] - '0'
		num02 := num2[index2] - '0'

		sums := int(num01) + int(num02) + left
		if sums >= 10 {
			left = 1
		} else {
			left = 0
		}

		tmp := (sums % 10) + '0'

		result = fmt.Sprintf("%c%s", tmp, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		num01 := num1[index1] - '0'
		sums := int(num01) + left
		if sums >= 10 {
			left = 1
		} else {
			left = 0
		}
		tmp := (sums % 10) + '0'
		result = fmt.Sprintf("%c%s", tmp, result)
		index1--
	}

	for index2 >= 0 {
		num02 := num2[index2] - '0'
		sums := int(num02) + left
		if sums >= 10 {
			left = 1
		} else {
			left = 0
		}
		tmp := (sums % 10) + '0'
		result = fmt.Sprintf("%c%s", tmp, result)
		index2--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}

	return result

}