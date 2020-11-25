package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
// 找出 nums 中的三个整数，使得它们的和与 target 最接近。
// 返回这三个数的和。假定每组输入只存在唯一答案。
//
//示例：
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 
//提示：
//3 <= nums.length <= 10^3
//-10^3 <= nums[i] <= 10^3
//-10^4 <= target <= 10^4

func threeSumClosest(nums []int, target int) int {
	len_nums := len(nums)
	closetNum := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	//fmt.Println(closetNum,nums)
	diff := math.Abs(float64(closetNum - target))
	for i := 0; i < len_nums-1; i++ {
		right := len_nums - 1
		left := i + 1
		for ; left < right; {
			result := nums[i] + nums[left] + nums[right]
			newDiff := math.Abs(float64(result - target))
			if newDiff < diff {
				diff = newDiff
				closetNum = result
			}
			if (result > target) {
				right -= 1
			} else if (result < target) {
				left += 1
			} else {
				return result
			}
		}
	}
	return closetNum
}

func main() {
	nums := []int{
		-1, 0, 0,-1,
	}
	s := threeSumClosest(nums, -100)
	fmt.Println(s, reflect.TypeOf(s))
}
