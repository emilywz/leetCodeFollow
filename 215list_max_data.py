"""
在未排序的数组中找到第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

说明:
你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
"""
class Solution:
    def findKthLargest(self, nums, k: int):
        if nums==[]:
            return None
        # tmp=nums
        # tmp.sort()
        tmp=self.sort_list(nums)
        print(tmp)
        return tmp[-k]


    def sort_list(self,nums):
        d_max=0
        d_min=0
        for i in nums:
            if d_max<i:
                d_max=i
            if d_min>i:
                d_min=i

        tmp={}
        for i in nums:
            if i in tmp.keys():
                tmp[i]+=1
            else:
                tmp[i]=1

        sort_list=[]
        for i in range(d_min,d_max+1):
            if i in tmp.keys():
                for j in range(tmp[i]):
                    sort_list.append(i)

        print(sort_list)
        return sort_list


if __name__ == '__main__':
    s=Solution()
    nums=[3,2,3,1,2,4,5,5,6]
    k=4
    target=s.findKthLargest(nums,k)
    print(target)
