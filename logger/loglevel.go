package logger

import "strings"

type logLevel int

const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
	CLOSE
)

func (l *logLevel)String()(lstr string) {
	switch *l {
	case 0: lstr = "DEBUG"
	case 1: lstr = "INFO"
	case 2: lstr = "WARNING"
	case 3: lstr = "ERROR"
	case 4: lstr = "FATAL"
	default:
		lstr = "CLOSE"
	}
	return
}

func Level(str string)(l logLevel) {
	str = strings.ToLower(str)
	switch str {
	case "debug":
		l = DEBUG
	case "info":
		l = INFO
	case "warning":
		l = WARNING
	case "error":
		l = ERROR
	case "fatal":
		l = FATAL
	default:
		l = CLOSE
	}
	return
}
