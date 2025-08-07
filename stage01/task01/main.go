package main

import "fmt"

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。 */

func main() {
	fmt.Printf("singleNumber([]int{2, 2, 1}): %v\n", singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	if len(nums) <= 0 {
		return 0
	} else if len(nums) <= 1 {
		return nums[0]
	}
	result := 0
	m := make(map[int]int, len(nums))
	for i := range nums {
		_, flag := m[nums[i]]
		if !flag {
			m[nums[i]] = nums[i]
		} else {
			delete(m, nums[i])
		}
	}

	for _, value := range m {
		result = value
	}
	return result
}
