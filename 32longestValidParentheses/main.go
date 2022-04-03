package main

import "fmt"

//32 最长有效括号
//
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0
//
//
//提示：
//
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	maxLen := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {

			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)

			} else {
				if maxLen < i-stack[len(stack)-1] {
					maxLen = i - stack[len(stack)-1]
				}

			}
		}

	}
	return maxLen

}

func main() {
	s := "(()"
	ret := longestValidParentheses(s)
	fmt.Println(ret)

}
