package main

import (
	"fmt"
	"sort"
)

// 多数元素
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//

// 示例 1：

// 输入：[3,2,3]
// 输出：3
// 示例 2：

// 输入：[2,2,1,1,1,2,2]
// 输出：2
//

// 进阶：

// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

func majorityElement(nums []int) int {

	sort.Ints(nums)
	return nums[len(nums)/2]

}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
