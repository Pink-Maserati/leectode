package main

import "fmt"

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串。

// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。

// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"

// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	left := 0
	right := 0
	for i := 0; i < n; i++ {
		dp[i][i] = true //从i到i肯定是回文
		for j := i - 1; j >= 0; j-- {
			//dp[i][j]表示从j到i的字符串是否是回文串
			dp[i][j] = s[j] == s[i] && (dp[j+1][i-1] || i-j == 1)
			if dp[i][j] && (i-j) > (right-left) {
				left = j
				right = i
			}

		}

	}
	return s[left : right+1]

}

//中心扩展法：我们枚举所有的「回文中心」并尝试「扩展」，直到无法扩展为止，此时的回文串长度即为此「回文中心」下的最长回文串长度。我们对所有的长度求出最大值，即可得到最终的答案。
func longestPalindrome1(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}

	}
	return s[start : end+1]

}
func expandAroundCenter(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome1(s))

}
