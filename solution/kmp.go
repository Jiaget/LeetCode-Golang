package solution

import "fmt"

// 对P字符串构建一个 前缀表。 （https://www.bilibili.com/video/BV1hW411a7ys/?spm_id_from=333.788.recommend_more_video.-1）
// 题目链接 （https://leetcode-cn.com/problems/implement-strstr/）
// 1. 第一步，对needle 做一个最大公共前缀表。

func prefixTable(needle string, prefix []int) {
	// 最大前缀表记录了到字符串的每一位止，出现过的公共前缀最长值。
	// 如abab, 前缀表为 0,0,1,2。

	n := len(needle)
	j := 0
	// 前缀表第一位永远是0
	prefix[0] = 0
	for i := 1; i < n; i++ {
		// 匹配失败，重置j
		for j > 0 && needle[j] != needle[i] {
			fmt.Println(j)
			j = prefix[j-1]
		}
		if needle[j] == needle[i] {
			// 匹配成功，公共前缀长度加一
			j++
			prefix[i] = j
		}
	}
}

func StrStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	prefix := make([]int, len(needle))
	prefixTable(needle, prefix)
	prefix = append([]int{-1}, prefix...)
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = prefix[j]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
