package main

import "fmt"

//31 下一个排列
//
//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须 原地 修改，只允许使用额外常数空间。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//示例 2：
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//示例 3：
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//示例 4：
//
//输入：nums = [1]
//输出：[1]
//
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 100

func nextPermutation(nums []int) {
	count := len(nums)

	if count <= 1 {

		return
	}

	i := count - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := count - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--

		}
		nums[i], nums[j] = nums[j], nums[i]

	}

	reverse(nums[i+1:])
}

func reverse(nums []int) {

	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}

}

func main() {
	num := []int{1, 1, 5}
	fmt.Println(&num)
	nextPermutation(num)

	fmt.Println(&num)
	a := 2
	b := 1
	a, b = b, a
	fmt.Println(a, b)

}
