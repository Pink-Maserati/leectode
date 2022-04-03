package main

import "fmt"

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"

// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

// 提示：
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成

//横向扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return prefix
		}

	}
	return prefix

}

func lcp(s1, s2 string) string {
	m := min(len(s1), len(s2))
	for i := 0; i < m; i++ {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	return s1[:m]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}

//纵向扫描
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]

			}
		}
	}
	return strs[0]
}

//分治
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft := lcp(start, mid)
		lcpRight := lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]

	}
	return lcp(0, len(strs)-1)
}
func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix1(strs))
	fmt.Println(longestCommonPrefix2(strs))
}
