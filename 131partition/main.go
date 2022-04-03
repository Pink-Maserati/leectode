package main

import "fmt"

// 131. 分割回文串
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

// 回文串 是正着读和反着读都一样的字符串。

// 示例 1：

// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：

// 输入：s = "a"
// 输出：[["a"]]

// 提示：
// 1 <= s.length <= 16
// s 仅由小写英文字母组成

func partition1(s string) (ans [][]string) {

	//回文判断 dp[i][j]表示s[i:j]是否是回文串 i<j
	//当j-i<1时，dp[i][j]=true
	//否则，dp[i][j]=dp[i+1][j-1] && s[i]==s[j]

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	//"abbab"输出 [["a","b","b","a","b"],["a","b","bab"],["a","bb","a","b"],["a","bbab"],["abba","b"]] 有问题
	// for i := 0; i < len(s); i++ {
	// 	for j := i + 1; j < len(s); j++ {
	// 		dp[i][j] = dp[i+1][j-1] && s[i] == s[j]

	// 	}
	// }

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]

		}
	}

	fmt.Println(dp)
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == len(s) {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < len(s); j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)

	return
}

func partition(s string) [][]string {
	res := [][]string{}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	fmt.Println(dp)
	dfs([]string{}, 0, &res, s, dp)
	return res
}

func dfs(res []string, start int, arr *[][]string, s string, dp [][]bool) {
	if start == len(s) {
		temp := make([]string, len(res))
		copy(temp, res)
		*arr = append(*arr, temp)

		return
	}

	for i := start; i < len(s); i++ {
		if dp[start][i] {
			res = append(res, s[start:i+1])

			dfs(res, i+1, arr, s, dp)

			res = res[:len(res)-1]
		}
	}

}

func partition3(s string) (arr [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	temp := []string{}
	var dfs func(int)

	dfs = func(i int) {
		if i == n {
			arr = append(arr, append([]string(nil), temp...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				temp = append(temp, s[i:j+1])
				dfs(j + 1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(0)

	return

}

func main() {

	s := "aab"
	//b := 'c'

	// b := [][]string{}
	// c := []string{"abc"}
	// b = append(b, c)
	// fmt.Println(b)

	arr := partition(s)
	fmt.Println(arr)
	arr1 := partition1(s)
	fmt.Println(arr1)
	arr2 := partition3(s)
	fmt.Println(arr2)

}
