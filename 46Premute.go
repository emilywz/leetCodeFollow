package main

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//[1,2,3],
//[1,3,2],
//[2,1,3],
//[2,3,1],
//[3,1,2],
//[3,2,1]
//]
//https://www.cnblogs.com/i-dandan/p/12443571.html
//回溯算法
func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	var tmp []int
	var visited = make([]bool, len(nums))
	listTrack(nums, &result, tmp, visited)
	return result
}
func listTrack(nums []int, result *[][]int, tmp []int, visited []bool) {
	//成功找到第一组
	if len(tmp) == len(nums) {
		var c = make([]int, len(tmp))
		copy(c, tmp)
		*result = append(*result, c)
		return
	}
	//回溯
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			tmp = append(tmp, nums[i])
			listTrack(nums, result, tmp, visited)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}

}
