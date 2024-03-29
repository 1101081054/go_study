package main

import (
	"unicode/utf8"
	"fmt"
)

var lastOccurred = make([]int, 0xffff)

/**
 * 最长不重复子串
 */
func lengthOfNonRepeatingSubStr(s string) int {
	//lastOccurred := make(map[rune]int)

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0;
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	str := "qwewertagj你好呀你好呀l"
	fmt.Println(len(str))	//字节长度
	fmt.Println(utf8.RuneCountInString(str))	//字符长度
	fmt.Println(lengthOfNonRepeatingSubStr(str))

}
