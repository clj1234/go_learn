package main

import (
	"fmt"
)

/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func main() {
	s := "{()}"
	fmt.Printf("isValid(s): %v\n", isValid(s))
}

func isValid(s string) bool {
	stack := []byte{}
	var m map[byte]byte = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range []byte(s) {
		if len(stack) > 0 && stack[len(stack)-1] == m[v] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) <= 0
}

func checkKeyExist(m map[byte]byte, b byte) bool {
	_, result := m[b]
	return result
}
