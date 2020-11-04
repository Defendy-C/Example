package redisopts

import (
	"MediaWeb/api"
	"MediaWeb/api/defs"
	"fmt"
	"strconv"
	"testing"
)

func Clear() {
	api.RedisWriteConn.Conn.Do("flushall")
}

func TestUserApi(t *testing.T)  {
	Clear()
	var u defs.User
	for i:=0;i < 10;i++ {
		u.Uid = i
		iStr := strconv.Itoa(i)
		u.Uname = "test" + iStr
		SetUserInfo(u)
	    fmt.Println(GetUserInfo(i))

	}
	Clear()
}