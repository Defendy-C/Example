package defs

import "time"

type User struct {
	Uid int `json:"uid"`
	Uname string `json:"uname"`
	Password string `json:"-"`
}

type VideoInfo struct {
	Vid string `json:"vid"`
	Uid int `json:"uid"`
	Vname string `json:"vname"`
	Vdate time.Time `json:"vdate"`
}

type CommentInfo struct {
	Cid string `json:"cid"`
	Uid int `json:"uid"`
	Uname string
	Vid string `json:"vid"`
	Content string `json:"content"`
	Cdate time.Time `json:"cdate"`
}
