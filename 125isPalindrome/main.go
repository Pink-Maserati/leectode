package main

import (
	"fmt"
	"strings"
)

// 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 解释："amanaplanacanalpanama" 是回文串
// 示例 2:

// 输入: "race a car"
// 输出: false
// 解释："raceacar" 不是回文串

// 提示：

// 1 <= s.length <= 2 * 105
// 字符串 s 由 ASCII 字符组成

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	start := 0
	end := len(s) - 1
	s = strings.ToUpper(s)

	// fmt.Printf("type %T, value %v\n", s[0], s[0])
	// fmt.Println(s)

	for start < end {
		for start < end && !isalnum(s[start]) {
			start++
		}
		for start < end && !isalnum(s[end]) {
			end--
		}
		if start < end {
			if s[start] != s[end] {
				return false
			}

			start++

			end--
		}

	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
