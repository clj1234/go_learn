package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

https://leetcode.cn/problems/two-sum/description/
*/
func main() {
	iarr := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("twoSum(iarr, target): %v\n", twoSum2(iarr, target))
}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	// map<num ,index>
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		value, flag := m[target-v]
		if flag && i != value {
			return []int{i, value}
		}
	}
	return []int{}
}
