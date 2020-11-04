package main

import (
	"strconv"
	"strings"
)

var b64map string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var b64pad byte = '='
var hexCode string = "0123456789abcdef"

//获取对应的16进制字符
func int2char(a int) byte {
	return hexCode[a]
}

//Base64转16进制
func b64tohex(s string) string {
	ret := ""
	k := 0
	slop := 0
	for i := 0;i < len(s);i++ {
		if s[i] == b64pad {
			break
		}
		var v int = strings.IndexByte(b64map, s[i])
		if v < 0 {
			continue
		}
		if k == 0 {
			ret += string(int2char(v >> 2))
			slop = v & 3
			k = 1
		} else if k == 1 {
			ret += string(int2char((slop << 2) | (v >> 4)))
			slop = v & 0xf
			k = 2
		} else if k == 2 {
			ret += string(int2char(slop))
			ret += string(int2char(v >> 2))
			slop = v & 3
			k = 3
		} else {
			ret += string(int2char((slop << 2) | (v >> 4)))
			ret += string(int2char(v & 0xf))
			k = 0
		}
	}
	if k == 1 {
		ret += string(int2char(slop << 2))
	}
	return ret
}

//16进制字符串转10进制
func parseInt(s string) int {
	temp_c, _ := strconv.ParseInt(s, 16, 0)
	c, _ := strconv.Atoi(strconv.FormatInt(temp_c, 10))
	return c
}

//16进制转Base64
func hex2b64(h string) string {
	var i, c int
	var ret string
	for i = 0;i + 3 <= len(h);i+=3 {
		c = parseInt(h[i : i + 3])
		ret += string(b64map[c >> 6])
		ret += string(b64map[c & 63])
	}
	if i + 1 == len(h) {
		c = parseInt(h[i : i + 1])
		ret += string(b64map[c << 2])
	} else if i + 2 == len(h) {
		c = parseInt(h[i : i + 2])
		ret += string(b64map[c >> 2])
		ret += string(b64map[(c & 3) << 4])
	}
	for ;(len(ret) & 3) > 0 ; {
		ret += string(b64pad)
	}
	return ret
}

