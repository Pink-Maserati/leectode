package main

import "fmt"

// 面试题 01.01. 判定字符是否唯一
// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

// 示例 1：

// 输入: s = "leetcode"
// 输出: false
// 示例 2：

// 输入: s = "abc"
// 输出: true
// 限制：

// 0 <= len(s) <= 100
// 如果你不使用额外的数据结构，会很加分。

func isUnique(astr string) bool {
	if len(astr) == 0 {
		return true
	}

	set := make(map[rune]bool, len(astr))
	for _, v := range astr {

		if set[v] {
			return false
		}

		set[v] = true
	}

	return true
}

func main() {
	s := "leetcode"

	fmt.Println(isUnique(s))
}
