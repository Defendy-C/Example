package dbopts

import (
	"MediaWeb/api"
	"MediaWeb/api/defs"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Clear() {
	conn := api.DBWriteConn.Conn
	conn.Exec("truncate user")
	conn.Exec("truncate video")
	conn.Exec("truncate session")
	conn.Exec("truncate comment")
}

func TestCommentAll(t *testing.T) {
	Clear()
	t.Run("Add", TestInsertComment)
	t.Run("Add", TestInsertUser)
	t.Run("Get", TestListComment)
	Clear()
}

func TestVideoAll(t *testing.T) {
	Clear()
	t.Run("Add", TestInsertVideoInfo)
	t.Run("Get", TestGetVideo)
	Clear()
}

func TestUserAll(t *testing.T) {
	Clear()
	t.Run("Add", TestInsertUser)
	t.Run("Get", TestGetUser)
	t.Run("Del", TestDelUser)
	Clear()
}

func TestInsertUser(t *testing.T) {
	var u defs.User
	u.Password = "test"
	for i := 0; i < 10;i++  {
		u.Uname = "test" + strconv.Itoa(i)
		err := InsertUser(u)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func TestDelUser(t *testing.T) {
	for i := 0; i < 10;i++  {
		uname := "test" + strconv.Itoa(i)
		err := DelUser(uname)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func TestGetUser(t *testing.T) {
	for i := 0;i < 10;i++ {
		uname := "test" + strconv.Itoa(i)
		u, err := GetUser(uname)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(u)
		}
	}
}

func TestInsertVideoInfo(t *testing.T) {
	var v defs.VideoInfo

	for i := 0;i < 10;i++ {
		v.Uid = i
		v.Vname = strconv.Itoa(i) + " video"
		err := InsertVideoInfo(v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestGetVideo(t *testing.T) {
	for i := 0; i < 10;i++  {
		v, err := GetVideoInfo(i)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println(v)
		}
	}
}

func TestInsertComment(t *testing.T) {
	for i:=0;i<10;i++ {
		istr := strconv.Itoa(i)
		c := defs.CommentInfo{Uid:i,Vid:istr,Content:"test"+istr}
		err := InsertComment(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestListComment(t *testing.T) {
	cs, err := ListComment("1", time.Date(2020,11,2,15,0,0,0,time.Local), time.Now())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*cs[0])
	}
}
