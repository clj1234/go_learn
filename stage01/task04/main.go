package main

import (
	"fmt"
)

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。

示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
*/
func main() {
	iarr := make([]int, 100)
	for i, _ := range iarr {
		iarr[i] = 9
	}
	// iarr := []int{1, 2, 3}
	fmt.Printf("plusOne(iarr): %v\n", plusOne(iarr))
}

// func plusOne(digits []int) []int {
// 	// 整型数组转为字符串数组
// 	strs := make([]string, len(digits))
// 	for i, v := range digits {
// 		strs[i] = strconv.Itoa(v)
// 	}
// 	// 字符串数组转为int64并++
// 	num, _ := strconv.ParseInt(strings.Join(strs, ""), 10, 64)
// 	num++
// 	// 拼接返回数组
// 	s := strconv.FormatInt(num, 10)
// 	iarr := make([]int, len(s))
// 	for i, v := range strings.Split(s, "") {
// 		iarr[i], _ = strconv.Atoi(v)
// 	}
// 	return iarr
// }

func plusOne(digits []int) []int {
	var result []int
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 && i == 0 {
			digits[i] = 0
			result = make([]int, len(digits)+1)
			copy(result[1:], digits)
			result[0] = 1
		} else if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			result = digits
			break
		}
	}
	return result
}
