package main

import "fmt"

// 搜索二维矩阵 II
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。

// 示例 1：

// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// 输出：true
// 示例 2：

// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// 输出：false

// 提示：

// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -109 <= matrix[i][j] <= 109
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -109 <= target <= 109

//从左下角或者右上角找
func searchMatrix(matrix [][]int, target int) bool {
	//列
	col := 0
	//行
	row := len(matrix) - 1
	// fmt.Println(col, row, len(matrix[0]))
	for row >= 0 && col <= len(matrix[0])-1 {

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			row--
		} else {
			col++
		}

	}

	return false

}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 20
	fmt.Println(searchMatrix(matrix, target))

}
