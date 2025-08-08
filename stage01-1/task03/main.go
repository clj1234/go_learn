package main

import (
	"fmt"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

https://leetcode-cn.com/problems/longest-common-prefix/
*/
func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Printf("longestCommonPrefix(s): %v\n", longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	result := strs[0]
	// 从第二位开始
	for i := 1; i < len(strs); i++ {
		tempResult := ""
		resultRunes := []rune(result)
		for j, v2 := range strs[i] {
			if len(resultRunes) > j && resultRunes[j] == v2 {
				tempResult = string(resultRunes[:j+1])
			} else {
				fmt.Println("v:", v2)
				break
			}
		}
		// 变短后重新赋值
		if len(tempResult) < len(result) {
			result = tempResult
		}
	}
	return result
}
