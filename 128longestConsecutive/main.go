package main

import (
	"fmt"
	"sort"
)

// 128. 最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9

// 提示：
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

//时间复杂度 o(nlogn)
func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	res := 1
	tmp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			tmp++

		} else {

			if tmp > res {
				res = tmp
			}

			tmp = 1

		}
		//fmt.Println(tmp, res)
	}
	if tmp > res {
		return tmp
	}
	return res
}

//时间复杂度：O(n)，其中 n 为数组的长度。具体分析已在上面正文中给出。
//空间复杂度：O(n) 。哈希表存储数组中所有的数需要 O(n) 的空间
func longestConsecutive1(nums []int) int {
	numSet := map[int]bool{}
	for _, v := range nums {
		numSet[v] = true
	}
	longRes := 0
	for k := range numSet {
		if !numSet[k-1] {
			cur := k
			tmp := 1
			for numSet[cur+1] {

				cur++
				tmp++
			}
			if tmp > longRes {
				longRes = tmp
			}
		}
	}
	return longRes
}

func main() {
	//	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive1(nums))
}
