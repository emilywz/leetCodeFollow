package main

import (
	"fmt"
	"strings"
	"unicode"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//输入: "race a car"
//输出: false

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	lowerS := strings.ToLower(s)
	lengthS := len(lowerS)
	lowerS2 := []rune(lowerS)
	left := 0
	right := lengthS - 1

	for left < right {
		letter_left := lowerS2[left]
		letter_right := lowerS2[right]
		isLeftLN := unicode.IsLetter(letter_left) || unicode.IsNumber(letter_left)
		isRightLN := unicode.IsLetter(letter_right) || unicode.IsNumber(letter_right)
		
		if !isLeftLN && isRightLN {
			left += 1
		}else if isLeftLN && !isRightLN {
			right -= 1
		}else if !isLeftLN && !isRightLN {
			left += 1
			right -= 1
		}else if letter_left==letter_right {
			left += 1
			right -= 1
		}else{
			return false
		}
	}
	return true
}

func main() {
	//s := "A man, a plan,   Panama"
	//s := "A man, a plan, a canal: Panama"
	s := ".,"
	//s := "race a car"
	fmt.Println(strings.ToLower(s))

	sb := isPalindrome(s)
	fmt.Println(sb)
	s1 := []rune(" ")
	fmt.Println(s1, string(s1), len(s1), unicode.IsLetter(s1[0]))
	//pattern := "\\d+" //反斜杠要转义
	//str := "124534"
	//result, _ := regexp.MatchString(pattern, str)
	//fmt.Println(result)
}
