package main

import "fmt"

// 347. 前 K 个高频元素
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]

// 示例 2:
// 输入: nums = [1], k = 1
// 输出: [1]

// 提示：

// 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

func topKFrequent(nums []int, k int) []int {
	// 	用map记录每个元素出现的次数
	// 按照出现次数放入List数组中
	// 放置的规则如下：
	// List数组中的下标是该元素出现的次数
	// List数组中的元素存放的是出现次数为该下标数值大小的数字
	// 最终从后往前（因为下标代表出现频率，频率高的在后面）取k个list数组中的元素
	numMap := map[int]int{}
	for _, v := range nums {
		numMap[v]++
	}
	//	fmt.Println(numMap)

	list := make([][]int, len(nums)+1)
	for k, v := range numMap {
		if list[v] == nil {
			list[v] = []int{}
		}
		list[v] = append(list[v], k)
	}
	//fmt.Println(list)
	res := make([]int, k)
	d := 0
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] != nil {
			tmp := list[i]
			for _, v := range tmp {
				if d < k {
					res[d] = v
					d++
				} else {
					break
				}
			}
		}
		if d >= k {
			break
		}
	}
	return res

}
func main() {
	//nums = [1,1,1,2,2,3], k = 2
	nums := []int{1, 1, 1, 2, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))

}
