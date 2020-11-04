package logger

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	time2 "time"
)

type FLog struct {
	 log *Log
	 maxSize int64
	 curFileSize int64
}

func NewFLog(l string, filePrefix string)*FLog {
	return &FLog{log: New(l, nil, filePrefix), maxSize:4096, curFileSize:0}
}

func (f *FLog)Write(l logLevel, text string)error {
	flag := false
	if f.log.writer == nil {
		err := f.newFile()
		fmt.Println(err)
		if err != nil {
			panic("flog.Write:" + err.Error())
		}
	} else {
		if f.curFileSize > f.maxSize {
			flag = true
			err := f.newFile()
			if err != nil {
				panic("flog.Write:" + err.Error())
			}
			f.curFileSize = 0
		}
	}
	f.curFileSize += int64(len(text))
	return f.log.write(l, text, flag)
}

func (f *FLog)MaxSize(n int64) {
	f.maxSize = n
}

func (f *FLog)newFile()error {
	time := time2.Now().Format("20060102")
	rnum := "" + strconv.Itoa(rand.Intn(999))
	filename := f.log.Prefix + time + rnum
	fmt.Println(filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	f.log.writer = file
	return nil
}

func (f *FLog)Debug(text string)error {
	return f.Write(DEBUG, text)
}

func (f *FLog)Info(text string)error {
	return f.Write(INFO, text)
}

func (f *FLog)Warning(text string)error {
	return f.Write(WARNING, text)
}

func (f *FLog)Error(text string)error {
	return f.Write(ERROR, text)
}

func (f *FLog)Fatal(text string)error {
	return f.Write(FATAL, text)
}
