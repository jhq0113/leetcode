package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

示例1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0

提示：

0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var (
		max   int
		start int
		list  = make(map[byte]int, 0)
	)

	for index := 0; index < len(s); index++ {
		char := s[index]
		if i, ok := list[char]; ok && i >= start {
			start = i + 1
			if index-i > max {
				max = index - i
			}
		} else if (index - start + 1) > max {
			max = index - start + 1
		}
		list[char] = index
	}

	return max
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		length = 1 + rand.Intn(15)
	)

	for start := 0; start < 10; start++ {
		testBytes := make([]byte, length, length)
		for index := 0; index < length; index++ {
			testBytes[index] = byte(97 + rand.Intn(25))
		}

		fmt.Println(string(testBytes))
		fmt.Println(lengthOfLongestSubstring(string(testBytes)))
	}
}
