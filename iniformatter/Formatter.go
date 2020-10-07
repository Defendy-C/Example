package iniformatter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Formatter struct {
	data []*DataPkg
}

type Err struct {
	lineNum int
	info string
}

func(e Err)Error()string {
	return fmt.Sprintf("At line %d: %s\n", e.lineNum, e.info)
}


// 检测是否为标识符
func check(str string)bool {
	ok, err := regexp.MatchString(`[_a-zA-Z][_a-zA-Z0-9]{0,30}`, str)
	if err != nil || !ok {
		return false
	}

	return true
}

func New(filename string)(f *Formatter, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	f = nil
	err = nil
	if err != nil {
		return
	}
	dp := make([]*DataPkg, 0)
	var dpTmp *DataPkg = nil
	var line string
	buf := bufio.NewReader(file)
	lineNum := 0
	for {
		line, err = buf.ReadString('\n')
		lineNum++
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		line = strings.Trim(line, "\r")
		line = strings.TrimSpace(line)
		// ignore empty line and note
		if line != "" && !strings.HasPrefix(line, "#") {
			length := len(line)
			// Add Tag
			if line[0] == '[' &&  line[length - 1] == ']' {
				var str string = line[1:length - 1]
				if !check(str) {
					err = Err{lineNum:lineNum, info: "not identified TagName"}
					break
				}
				if dpTmp != nil {
					dp = append(dp, dpTmp)
				}
				dpTmp = NewDataPkg(str)
			} else {
				// Add Attribute
				ss := strings.Split(line, "=")
				if len(ss) != 2 && len(ss) != 1 {
					err = Err{lineNum:lineNum, info: "unidentified key-value"}
					break
				}
				key := strings.TrimSpace(ss[0])
				var val string
				if len(ss) == 1 {
					val = "true"
				} else {
					val = strings.TrimSpace(ss[1])
				}
				dpTmp.attr[key] = val
			}
		}
	}
	dp = append(dp, dpTmp)
	f = &Formatter{data:dp}
	return
}

func (f *Formatter)Print() {
	for _, val := range f.data {
		fmt.Printf("[%s]\n", val.pkgName)
		for k, v := range val.attr {
			fmt.Printf("  %s : %s\n", k, v)
		}
	}
}

func (f *Formatter)GetDataPkgByName(name string) *DataPkg {
	for _, v := range f.data {
		if v.pkgName == name {
			return v
		}
	}
	return nil
}

