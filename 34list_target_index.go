package main

import "fmt"

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。
//找出给定目标值在数组中的开始位置和结束位置。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//如果数组中不存在目标值，返回 [-1, -1]。
//
//示例 1:
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//
//示例 2:
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]

func searchRange(nums []int, target int) []int {
	length := len(nums)
	not_exist := []int{-1, -1}
	if length == 0 {
		return not_exist
	}
	if target > nums[length-1] || target < nums[0] {
		return not_exist
	}
	var result, result_index []int
	first := BinarySearchFirst(nums, target)
	if first == -1 {
		return not_exist
	}
	for i := first; i < length; i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}
	totalCount := len(result) - 1
	if totalCount == 0 {
		result_index = append(result_index, result[0], result[0])
	} else {
		result_index = append(result_index, result[0], result[totalCount])
	}
	return result_index
}
//二分查找第一个目标值
func BinarySearchFirst(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func sortList(nums []int) []int {
	if len(nums) <= 1 {
		return nil
	}
	min, max := countMinMax(nums)

	tmp := countSort(max, nums)
	//fmt.Println("tmp",tmp)

	var index int
	for i := min; i < len(tmp); i++ {
		for j := tmp[i]; j > 0; j-- {
			nums[index] = i
			index++
		}
	}
	return nums
}

func countSort(max int, nums []int) []int {
	tmp := make([]int, max+1)
	for i := 0; i < len(nums); i++ {
		tmp[nums[i]] += 1
		fmt.Println("i", i, "tmp", tmp[nums[i]], "num", nums[i])
	}
	return tmp
}

func countMinMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min, max
}

func main() {
	//results := []int{
	//	5, 7, 7, 8, 8, 10,
	//}
	var result []int
	results := []int{
	 6, 2, 1, 7, 4, 5, 9, 8, 3,
	}
	fmt.Println(results)
	//min, max := countMinMax(results)
	//tmp := countSort(max, results)
	sortList(results)
	//for _, d := range results{
	//	fmt.Printf("%d ", d)
	//}
	//result = searchRange(results, 5)
	//fmt.Println(results, "tmp", tmp, "min", min, "max", max, "result", result)
	fmt.Println(results, "result", result)
}
