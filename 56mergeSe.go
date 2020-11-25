package main

import (
	"fmt"
	"sort"
)

//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//示例 2:
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

type MergeSlice [][]int

func (i MergeSlice) Len() int {
	return len(i)
}

func (i MergeSlice) Swap(n, m int) {
	i[n], i[m] = i[m], i[n]
}

func (i MergeSlice) Less(n, m int) bool {
	return i[n][0] < i[m][0]
}

func max(n,m int) int  {
	if n>=m {
		return n
	}else{
		return m
	}
}

func merge(intervals [][]int) [][]int {
	lenInterval := len(intervals)
	if lenInterval <= 0 {
		return nil
	}
	sort.Sort(MergeSlice(intervals))
	//fmt.Println("sort [][]",intervals)
	tmpList:=make([][]int,0)

	var i int
	for i<lenInterval{
		left:=intervals[i][0]
		right:=intervals[i][1]
		j:=i+1
		//fmt.Println("data [][]:",left,right,j)
		for j<lenInterval{
			if intervals[j][0]<=right{
				right=max(intervals[j][1],right)
				j++
			}else{
					break
			}
		}
		tmp:=[]int{left,right}
		tmpList=append(tmpList,tmp)
		//fmt.Println("right:",right,j,tmpList)
		i=j
	}
	return tmpList
}
