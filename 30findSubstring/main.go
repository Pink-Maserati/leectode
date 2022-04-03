package main

import "fmt"

//30 串联所有单词的子串
//给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//
//注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。
//
//
//
//示例 1：
//
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
//示例 2：
//
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//示例 3：
//
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]
//
//
//提示：
//
//1 <= s.length <= 104
//s 由小写英文字母组成
//1 <= words.length <= 5000
//1 <= words[i].length <= 30
//words[i] 由小写英文字母组成

func findSubstring(s string, words []string) []int {
	var res []int
	strlen := len(s)

	wordsLen := len(words)

	if strlen == 0 || wordsLen == 0 {
		return res
	}
	oneLen := len(words[0])
	if strlen < wordsLen*oneLen {
		return res
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for left := 0; left < strlen-wordsLen*oneLen+1; left++ {
		flag := true
		tmpMap := make(map[string]int)
		for right := left; right < left+wordsLen*oneLen; right += oneLen {
			tmp := s[right : right+oneLen]
			if wordMap[tmp] == 0 {
				flag = false
				break
			}
			tmpMap[tmp]++
			if wordMap[tmp] < tmpMap[tmp] {
				//单词出现次数太多

				flag = false
				break
			}

		}
		if flag {
			res = append(res, left)
		}

	}

	return res

}
func findSubstring1(s string, words []string) []int {
	// 获取单词长度
	wordLength := len(words[0])
	wordNum := len(words)
	stringLenth := len(s)
	var ans []int
	// 没有单词
	if wordNum == 0 {
		return ans
	}
	if stringLenth < wordLength {
		return ans
	}
	// 各单词应出现数量
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	// 构建滑动窗口
	for left := 0; left <= stringLenth-wordLength*wordNum; left++ {
		// 创建一个map记录该滑动窗口单词出现情况
		tempMap := make(map[string]int)
		flag := true // 记录子串是否合法
		for right := left; right < left+wordLength*wordNum; right += wordLength {
			// 截取一个单词
			temp := s[right : right+wordLength]
			if wordMap[temp] == 0 {
				// 没有这个单词
				flag = false
				break
			}
			// 有这个单词，记录到map中
			tempMap[temp]++
			if wordMap[temp] < tempMap[temp] {
				// 单词出现次数太多（太少可能在后面会出现）
				flag = false
				break
			}
		}
		if flag {
			// 通过检验
			ans = append(ans, left)
		}
	}
	return ans
}

func main() {
	s := "aaa"
	// fmt.Println(s[0:2])
	for i := range s {
		fmt.Println(i)
	}
	words := make([]string, 0)
	words = append(words, "a")

	words = append(words, "a")

	fmt.Println(findSubstring(s, words))
	fmt.Println(findSubstring1(s, words))

}
