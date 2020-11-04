package common

import (
	"MediaWeb/tools/logger"
	"time"
)

var (
	Path string           = "C:\\Users\\Defendy-C\\Documents\\goworkplace\\src\\MediaWeb"
	TimeOut time.Duration = time.Hour * 5
	Log *logger.FLog      = logger.NewFLog("DEBUG", Path + "\\log\\flog")

)

func init() {
	Log.MaxSize(1024*1024*10)
}
