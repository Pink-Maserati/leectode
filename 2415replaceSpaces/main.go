package main

import (
	"fmt"
	"strings"
)

// 面试题 01.03. URL化
// URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

// 示例 1：

// 输入："Mr John Smith    ", 13
// 输出："Mr%20John%20Smith"
// 示例 2：

// 输入："               ", 5
// 输出："%20%20%20%20%20"

// 提示：
// 字符串长度在 [0, 500000] 范围内。

func replaceSpaces(S string, length int) string {

	if length == 0 || len(S) == 0 {
		return ""
	}

	ret := ""
	for i := 0; i < length; i++ {
		if S[i] != ' ' {
			ret = ret + string(S[i])

		} else {
			ret = ret + "%20"
		}
	}
	return ret
}

//性能优于第一种
func replaceSpaces1(S string, length int) string {
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		if S[i] != ' ' {
			builder.WriteByte(S[i])
		} else {
			builder.WriteString("%20")
		}
	}
	return builder.String()
}

func main() {
	s := "Mr John Smith    "

	len := 13
	ret := replaceSpaces(s, len)
	fmt.Println(ret)
	ret1 := replaceSpaces1(s, len)
	fmt.Println(ret1)

}
