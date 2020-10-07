package unittest

import "strings"

// make一个切片 alloc 1次
// 切片赋值给 strs alloc 1次
// 创建返回值切片 alloc 1次
func SplitString(src string, delim string) []string {
	strs := make([]string, strings.Count(src, delim)+1)
	delimLen := len(delim)
	index := strings.Index(src, delim)
	i := 0
	for index >= 0 {
		strs[i] = src[:index]
		i++
		src = src[index+delimLen:]
		index = strings.Index(src, delim)
	}
	strs[i] = src

	return strs
}
