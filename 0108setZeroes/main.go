package main

import (
	"fmt"
)

// 面试题 01.08. 零矩阵
// 编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

// 示例 1：
// 输入：
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// 输出：
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]

// 示例 2：
// 输入：
// [
//   [0,1,2,0],
//   [3,4,5,2],
//   [1,3,1,5]
// ]
// 输出：
// [
//   [0,0,0,0],
//   [0,4,5,0],
//   [0,3,1,0]
// ]
func setZeroes(matrix [][]int) {

	col := make([]bool, len(matrix[0]))
	row := make([]bool, len(matrix))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				col[j] = true
				row[i] = true

			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if col[j] || row[i] {
				r[j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)
}
