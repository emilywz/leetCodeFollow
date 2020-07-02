package main

import "fmt"

//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
//
//示例 1:
//输入:
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出: 3
//解释:
//长度最长的公共子数组是 [3, 2, 1]。
//说明:
//1 <= len(A), len(B) <= 1000
//0 <= A[i], B[i] < 100

func findLength(A []int, B []int) int {
	lengthA := len(A)
	lengthB := len(B)
	if lengthA == 0 || lengthB == 0 {
		return 0
	}
	maxLength := 0
	end := 0

	rows := 0
	cols := lengthB - 1
	for rows < cols {
		i, j := rows, cols
		length := 0
		for i < lengthA && j < lengthB {
			if A[i] != B[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if cols > 0 {
			cols--
		} else {
			rows++
		}
	}
	result := A[(end - maxLength + 1):(end + 1)]
	fmt.Println("result",result)
	return len(result)
}

//不连续的最长公约子串长度
func findLength(A []int, B []int) int {
	lengthA := len(A)
	lengthB := len(B)
	if lengthA == 0 || lengthB == 0 {
		return 0
	}
	lookup := make([][]int, lengthA+1)
	//for i:=range lookup{
	//	lookup[i] = make([]int, lengthB+1)
	//}
	for i:=0;i<len(lookup);i++ {
		lookup[i] = make([]int, lengthB+1)
	}

	var i, j int
	//for i = 0; i <= lengthA; i++ {
	//	lookup[i][0] = 0
	//}
	//for j = 0; i <= lengthB; j++ {
	//	lookup[0][j] = 0
	//}
	result:=0
	for i = lengthA-1; i >=0; i-- {
		for j = lengthB-1; j >=0; j-- {
			if A[i] == B[j] {
				lookup[i][j] = lookup[i+1][j+1] + 1
			} else {
				//len1 := lookup[i-1][j]
				//len2 := lookup[i][j-1]
				//if len1 > len2 {
				//	lookup[i][j] = len1
				//} else {
				//	lookup[i][j] = len2
				//}
				lookup[i][j]=0
			}
			if result<lookup[i][j]{
				result=lookup[i][j]
			}
		}
	}
	fmt.Println("lookup",lookup)
	return result
}

func main() {
	A:=[]int{
		1,2,3,2,1,4,
		//0,1,1,1,1,
	}
	B:=[]int{
		3,2,1,4,7,
		//1,0,1,0,1,
	}
	s:=findLength(A,B)
	fmt.Println("len",s)

	//r:=findUncontinueLength(A,B)
	//fmt.Println("unlen",r)
}
