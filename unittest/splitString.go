package unittest

import "strings"

func SplitString(src string, delim string)[]string {
	strs := make([]string, 0)
	delimLen := len(delim)
	index := strings.Index(src, delim)
	for index >= 0 {
		strs = append(strs, src[:index])
		src = src[index+delimLen :]
		index = strings.Index(src, delim)
	}
	strs = append(strs, src)

	return strs
}
