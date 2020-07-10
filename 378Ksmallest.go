package main

import (
	"fmt"
	"sort"
)

//给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
//

//示例：
//matrix = [
//[ 1,  5,  9],
//[10, 11, 13],
//[12, 13, 15]
//],
//k = 8,
//返回 13。

//提示：
//你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。

func kthSmallest(matrix [][]int, k int) int {
	lengthM := len(matrix)
	if lengthM == 0 {
		return -1
	}
	//min := matrix[0][0]
	//max := matrix[lengthM-1][lengthM-1]
	tmp := make(map[int]int, 0)
	var nums,sortNums []int
	//index :=  0
	for i := 0; i < lengthM; i++ {
		for j := 0; j < lengthM; j++ {
			if _, ok := tmp[matrix[i][j]]; ok {
				tmp[matrix[i][j]] +=1
			} else {
				tmp[matrix[i][j]] = 1
			}
			//fmt.Println("i", i, "tmp", tmp[matrix[i][j]], "num", matrix[i][j])
			nums = append(nums, matrix[i][j])
		}
	}

	//var index int
	fmt.Println(tmp)
	//for i := min; i < max+1; i++ {
	//	if _, ok := tmp[i]; ok {
	//		for j := tmp[i]; j > 0; j-- {
	//			sortNums = append(sortNums, i)
	//			fmt.Println("i", i, "nums index", index)
	//			index++
	//		}
	//	} else {
	//		continue
	//	}
	//}

	sort.Ints(nums)
	fmt.Println(sortNums)
	result := nums[k-1]
	return result

}

func main() {
	ms := [][]int{
		{4, 4, 8}, {4, 4, 9}, {5, 6, 13},
	}
	data:=kthSmallest(ms,8)
	fmt.Println(data)
}
