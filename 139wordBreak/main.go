package main

import "fmt"

// 139.单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

//

// 示例 1：s

// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
// 示例 2：

// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//      注意，你可以重复使用字典中的单词。
// 示例 3：

// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false
//

// 提示：

// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中的所有字符串 互不相同

func wordBreak(s string, wordDict []string) bool {

	//dp[i]：表示s从0~i-1即s[0:i-1]字串能否被拆分成单词
	//dp[0]:true
	//dp[i]=dp[j]&&check(s[j:i])，check(s[j:i])检查字串s[j:i-1]是否出现在字典中

	dp := make([]bool, len(s)+1)
	//用于判断单词在不在字典中
	wordDictMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordDictMap[v] = true
	}

	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}

func main() {
	// s := "catsandog"
	// wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))

	s = "catsandog"

	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))

}
