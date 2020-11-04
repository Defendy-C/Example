package controller

import (
	"MediaWeb/api/dbopts"
	"MediaWeb/api/defs"
	"MediaWeb/common"
	"fmt"
	"net/http"
	"strconv"
)

func DealUser(w http.ResponseWriter ,r *http.Request) {
	// 路径检测
	var dir []string
	dir = urlToSliceAndCheck(r.URL.Path[1:])
	if len(dir) != 1 {
		ErrResponse(&w, defs.ErrNotPage)
		return
	}

	// 获取post数据
	r.ParseForm()
	data := r.PostForm
	u := &defs.User{Uname:data.Get("uname"), Password:data.Get("password")}

	// 验证token
	tokenStr := r.Header.Get("X-Token")
	var tokenIsValid int
	var err error
	if !(dir[0] == "add" || dir[0] == "login" && tokenStr == "") {
		tokenIsValid, _, err = verifyToken(u.Uname, tokenStr)
		if err != nil {
			ErrResponse(&w, defs.ErrTokenVerify)
			return
		}
	}

	if u.Uname == "" || u.Password == "" {
		ErrResponse(&w, defs.ErrParseRequest)
		return
	} else {
		var dbErr error
		var successStr string
		switch dir[0] {
		case "add" :successStr, dbErr = registerUser(*u)
		case "login" :
			successStr, dbErr = loginUser(tokenIsValid, u)
			if tokenIsValid == 0 {
				tokenStr = refreshToken(*u, &w)
			}
			tokenIsValid = 1
		case "exit" :
			clearToken(&w)
			tokenIsValid = 0
		default:
			ErrResponse(&w, defs.ErrNotPage)
			return
		}

		if tokenIsValid == 2 {
			tokenStr = refreshToken(*u, &w)
		}
		successStr += ",token=" + tokenStr
		dealDbStatus(&w, successStr, dbErr)
	}
}

func DealVideo(w http.ResponseWriter, r *http.Request) {
	// 路径检测
	var dir []string
	dir = urlToSliceAndCheck(r.URL.Path[1:])
	if len(dir) < 2 {
		ErrResponse(&w, defs.ErrNotPage)
	}

	// 获取post数据
	r.ParseForm()
	data := r.PostForm
	v := &defs.VideoInfo{Vname:data.Get("vname")}
	if v.Vname == "" {
		ErrResponse(&w, defs.ErrParseRequest)
	}

	// 验证token
	tokenStr := r.Header.Get("X-Token")
	id, err := strconv.Atoi(dir[1])
	if err != nil {
		panic(err)
		common.Log.Error(err.Error())
	}
	uid, tokenIsValid, err := verifyToken(getSecret(id), tokenStr)
	if tokenIsValid == 0 || err != nil || uid != id {
		fmt.Println(tokenIsValid, err, id, uid)
		ErrResponse(&w, defs.ErrTokenVerify)
		return
	}

	switch dir[0] {
	case "add":dbopts.InsertVideoInfo(*v)
	}


}
