package main

import (
	"fmt"
	"strconv"
)

// 241. 为运算表达式设计优先级
// 给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。

// 示例 1：
// 输入：expression = "2-1-1"
// 输出：[0,2]
// 解释：
// ((2-1)-1) = 0
// (2-(1-1)) = 2

// 示例 2：
// 输入：expression = "2*3-4*5"
// 输出：[-34,-14,-10,-10,10]
// 解释：
// (2*(3-(4*5))) = -34
// ((2*3)-(4*5)) = -14
// ((2*(3-4))*5) = -10
// (2*((3-4)*5)) = -10
// (((2*3)-4)*5) = 10

// 提示：
// 1 <= expression.length <= 20
// expression 由数字和算符 '+'、'-' 和 '*' 组成。
// 输入表达式中的所有整数值在范围 [0, 99]
func diffWaysToCompute(expression string) []int {
	if len(expression) == 0 {
		return []int{0}
	}

	//如果是数字，直接返回
	if isDigit(expression) {
		tmp, _ := strconv.Atoi(expression)
		return []int{tmp}
	}

	res := []int{}
	for i, v := range expression {
		if v == '+' || v == '-' || v == '*' {
			//如果是运算符，则计算左右两边的算式
			leftNums := diffWaysToCompute(expression[:i])
			rightNums := diffWaysToCompute(expression[i+1:])
			for _, l := range leftNums {
				for _, r := range rightNums {
					tmp := 0
					if v == '+' {
						tmp = l + r
					} else if v == '-' {
						tmp = l - r
					} else if v == '*' {
						tmp = l * r
					}
					res = append(res, tmp)
				}
			}
		}
	}
	return res

}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)

	return err == nil
}

func main() {
	// expression := "2-1-1"
	expression := "2*3-4*5"
	fmt.Println(diffWaysToCompute(expression))
}
