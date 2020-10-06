package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Log struct {
	ch chan *logInfo
	writer io.Writer
	level logLevel
	Prefix string
}

type logInfo struct {
	flag bool  // 用于切换写入源
	writer io.Writer
	text string

}

func New(l string, w io.Writer, prefix string)*Log {
	ch := make(chan *logInfo, 128)
	level := Level(l)
	logObj := Log{level:level, Prefix:prefix, ch:ch, writer:w}
	go logObj.syncPrint()
	return &logObj
}

func(l *Log)packLog(text string) []byte {
	var tbs []byte
	nowTime := time.Now().Format("2006/01/02 15:03:04")
	preStr := fmt.Sprintf("%s %s [%s] : ", l.Prefix, nowTime, l.level.String())
	tbs = append(tbs, preStr...)
	tbs = append(tbs, text...)
	tbs = append(tbs, '\n')
	return tbs
}

func(l *Log)syncPrint() {
	for v := range l.ch {
		if v.writer == nil {
			fmt.Println("log.syncPrint: writer is nil, don't print log \"", v.text, "\"")
			continue
		}

		tbs := l.packLog(v.text)
		w := bufio.NewWriter(v.writer)
		length := 0
		size := len(tbs)
		for {
			n, err := w.Write(tbs)
			length += n
			if err != nil {
				fmt.Println(err)
				l.write(ERROR, err.Error(), v.flag)
				v.flag = false
				break
			}
			if length >= size {
				break
			}
		}
		w.Flush()
		if v.flag {
			f, ok := v.writer.(*os.File)
			if ok {
				f.Close()
			}
		}
	}
}

func(l *Log)enable(level logLevel) bool {
	if level < l.level {
		return false
	}
	return true
}

func (l *Log)write(level logLevel, text string, close bool)error { // close 是否写完当前数据后关闭writer
	info := logInfo{flag:close, text:text, writer:l.writer}
	if l.enable(level) {
		l.ch <- &info
	}

	return nil
}

func (l *Log)Debug(text string)error {
	return l.write(DEBUG, text, false)
}

func (l *Log)Info(text string)error {
	return l.write(INFO, text, false)
}

func (l *Log)Warning(text string)error {
	return l.write(WARNING, text, false)
}

func (l *Log)Error(text string)error {
	return l.write(ERROR, text, false)
}

func (l *Log)Fatal(text string)error {
	return l.write(FATAL, text, false)
}

func (l *Log)Writer(w io.Writer) {
	l.writer = w
}
