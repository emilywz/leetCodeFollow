package main

import (
	"fmt"
	"math"
)

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//输入: "cbbd"
//输出: "bb"

func longestPalindrome1(s string) string {
	lengthS := len(s)
	if lengthS == 0 || lengthS == 1 {
		return s
	}
	midNum := lengthS / 2
	var start, end int
	leftS := s[:midNum]
	lengthLeft := len(leftS)

	var rightStart int

	if lengthS%2 == 0 {
		rightStart = midNum
	} else {
		//midS:=s[midNum]
		rightStart = midNum + 1
	}
	rightS := s[rightStart:]
	for i := 1; i <= midNum; i++ {
		if rightS[:i] == leftS[lengthLeft-i:] {
			start = midNum - i
			end = rightStart + i
		}
	}
	return s[start:end]
}
func longestPalindrome2(s string) string {
	lengthS := len(s)
	if lengthS <= 1 {
		return s
	}
	midNum := lengthS / 2
	var start, end int = 0, 1
	//leftS := s[:midNum]
	//lengthLeft := len(leftS)
	//var rightStart int
	//if lengthS%2 == 0 {
	//	rightStart = midNum
	//} else {
	//	//midS:=s[midNum]
	//	rightStart = midNum + 1
	//}
	//rightS := s[rightStart:]
	//for i := 1; i <= midNum; i++ {
	//	if rightS[:i] == leftS[lengthLeft-i:] {
	//		start = midNum - i
	//		end = rightStart + i
	//	}
	//}

	for i := 0; i < lengthS; {
		if lengthS-i <= midNum {
			break
		}

		b, e := i, i
		for e < lengthS-1 && s[e+1] == s[e] {
			e++
		}
		for e < lengthS-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > end {
			start = b
			end = newLen
		}
	}
	return s[start : start+end]
}

func longestPalindrome3(s string) string {
	lengthS := len(s)
	if lengthS <= 1 {
		return s
	}
	//midNum := lengthS / 2
	//var start, end int = 0, 1
	longest := s[0:1]
	for i := 0; i < lengthS; i++ {
		for rightStep := 0; rightStep < 2; rightStep++ {
			for p, q := i-1, i+rightStep; p >= 0 && q < lengthS && s[p] == s[q]; {
				if q-p+1 > len(longest) {
					longest = s[p : q+1]
				}
				p--
				q++
			}
		}
	}
	return longest
}

func longestPalindrome(s string) string {
	lengthS := len(s)
	if lengthS <= 1 {
		return s
	}
	//midNum := lengthS / 2
	var start, end int = 0, 0

	for i := 0; i < lengthS; i++ {
		len1:=expandAroundCenter(s,i,i)
		len2:=expandAroundCenter(s,i,i+1)
		lens:=int(math.Max(float64(len1),float64(len2)))
		if lens>end-start{
			start=i-(lens-1)/2
			end=i+lens/2
		}
	}
	return string(s[start:end+1])
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R-L-1
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	s := "bababd"
	re := longestPalindrome(s)
	fmt.Println(re)
}
