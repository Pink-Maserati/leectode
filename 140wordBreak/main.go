package main

import (
	"fmt"

	"strings"
)

// 140. 单词拆分 II
// 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。

// 注意：词典中的同一个单词可能在分段中被重复使用多次。

// 示例 1：

// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
// 示例 2：

// 输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
// 输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
// 解释: 注意你可以重复使用字典中的单词。
// 示例 3：

// 输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// 输出:[]

// 提示：

// 1 <= s.length <= 20
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 10
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中所有字符串都 不同

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	var arr []string
	dfs(s, wordSet, []string{}, &arr)

	return arr
}

func dfs(s string, wordSet map[string]bool, temp []string, res *[]string) {
	if len(s) < 1 { // 字符串遍历结束，说明完全匹配，拼接后写入结果
		*res = append(*res, strings.Join(temp, " "))
		return

	}
	for i := 0; i < len(s)+1; i++ { //逐一作为切分点
		if wordSet[s[:i]] {
			t := make([]string, len(temp))
			copy(t, temp)
			t = append(t, s[:i]) //// 记录当前结果进入递归
			dfs(s[i:], wordSet, t, res)
		}
	}

}

func wordBreak1(s string, wordDict []string) (res []string) {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了93.53%的用户
	wordMap := map[string]bool{} // 哈希表加快访问速度
	for _, word := range wordDict {
		wordMap[word] = true
	}
	check(s, &[]string{}, &res, &wordMap)
	return

}

func check(s string, tempSlice *[]string, res *[]string, wordMap *map[string]bool) {
	if len(s) < 1 { // 字符串遍历结束，说明完全匹配，拼接后写入结果
		*res = append(*res, strings.Join(*tempSlice, " "))
		return
	}

	for i := 0; i < len(s)+1; i++ { // 逐一作为切分点
		if (*wordMap)[s[:i]] { // 当前切分结果在字典中则继续
			// 做拷贝
			temp := make([]string, len(*tempSlice))
			copy(temp, *tempSlice)
			temp = append(temp, s[:i]) // 记录当前结果进入递归
			check(s[i:], &temp, res, wordMap)
		}
	}
}



func main() {
	// a := "hahaha"
	// b := "hehehe"
	// c := strings.Join([]string{a, b}, ",")
	// println(c)
	s := "pineapplepenapple"
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	arr := wordBreak(s, wordDict)
	fmt.Println("arr:")
	for _, v := range arr {
		fmt.Println(v)
	}
	arr1 := wordBreak1(s, wordDict)
	fmt.Println("arr1:")
	for _, v := range arr1 {
		fmt.Println(v)
	}


}
