package main

import (
	"fmt"
)

// 287. 寻找重复数
// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

// 示例 1：
// 输入：nums = [1,3,4,2,2]
// 输出：2

// 示例 2：
// 输入：nums = [3,1,3,4,2]
// 输出：3

// 提示：
// 1 <= n <= 105
// nums.length == n + 1
// 1 <= nums[i] <= n
// nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

// 进阶：
// 如何证明 nums 中至少存在一个重复的数字?
// 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？

//时间复杂度：O(n)。「Floyd 判圈算法」时间复杂度为线性的时间复杂度。
//空间复杂度：O(1)。我们只需要常数空间存放若干变量。
func findDuplicate(nums []int) int {

	//[1,3,4,2]:
	// 0->1
	// 1->3
	// 2->4
	// 3->2
	//0->1->3->2->4->null
	//[1,3,4,2,2]:
	// 	0->1
	// 1->3
	// 2->4
	// 3->2
	// 4->2
	//0->1->3->2->4->2->4->2->……
	//slow.Next===>nums[slow]
	//fast.Next.Next====>nums[nums[fast]]
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	p := 0
	for p != slow {
		slow = nums[slow]
		p = nums[p]
	}
	return p

}

//应用一个性质: 如果count(nums, mid) <= mid则该数未重复, 进而使用二分法。
//时间复杂度:O(nlogn)，其中 n 为nums 数组的长度。
//二分查找最多需要二分 O(logn) 次，每次判断的时候需要O(n) 遍历 nums 数组求解小于等于 mid 的数的个数，因此总时间复杂度为O(nlogn)。
//空间复杂度：O(1)。我们只需要常数空间存放若干变量
func findDuplicate1(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		// 表明当前数未重复
		if count(nums, mid) <= mid {
			left = mid + 1

		} else {
			//当前mid可能是缺失的, 也可能不是, 继续向左侧压缩搜索区间。
			right = mid
		}
	}
	return left

}

func findDuplicate2(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

/// 查找数组中比mid小的数的个数
func count(nums []int, mid int) int {
	n := 0
	for i := range nums {
		if nums[i] <= mid {
			n++
		}
	}
	return n
}

func main() {
	//nums = [1,3,4,2,2]
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
	fmt.Println(findDuplicate1(nums))
	fmt.Println(findDuplicate2(nums))

}
