package main

import (
	"fmt"
	"strings"
)

// 151. 翻转字符串里的单词
// 给你一个字符串 s ，逐个翻转字符串中的所有 单词 。

// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。

// 说明：

// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。

// 示例 1：
// 输入：s = "the sky is blue"
// 输出："blue is sky the"

// 示例 2：
// 输入：s = "  hello world  "
// 输出："world hello"
// 解释：输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包括。

// 示例 3：
// 输入：s = "a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，将翻转后单词间的空格减少到只含一个。

// 示例 4：
// 输入：s = "  Bob    Loves  Alice   "
// 输出："Alice Loves Bob"

// 示例 5：
// 输入：s = "Alice does not even like bob"
// 输出："bob like even not does Alice"

// 提示：
// 1 <= s.length <= 104
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词

// 进阶：
// 请尝试使用 O(1) 额外空间复杂度的原地解法。

func reverseWords(s string) string {
	//去除前后的空格
	s = strings.TrimSpace(s)
	//分割字符串
	s1 := strings.Split(s, " ")

	var res strings.Builder
	//反转拼接
	for i := len(s1) - 1; i >= 0; i-- {
		//去掉空间的空字符串
		s2 := strings.TrimSpace(s1[i])
		if len(s2) == 0 {
			continue
		}
		res.WriteString(s2)
		if i != 0 {
			res.WriteString(" ")
		}

	}

	return res.String()
}

//栈：利用栈的先进后出的特性，实现字符串的翻转
func reverseWords1(s string) string {
	left := 0
	right := len(s) - 1
	stack := []string{}
	var str strings.Builder
	for left <= right {
		//过滤空格
		if s[left] != ' ' {
			str.WriteByte(s[left])
			//单词前面是空格 或者是 最后一个单词就入栈
			if left == right || s[left+1] == ' ' {
				stack = append(stack, str.String())
				str = strings.Builder{}

			}
		}
		left++
	}

	var res strings.Builder
	for len(stack) > 0 {
		res.WriteString(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		if len(stack) != 0 {
			res.WriteString(" ")
		}
	}
	return res.String()

}

func main() {
	// s := "  Alice does not even like bob"
	s := "a good   example"

	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords1(s))

}
