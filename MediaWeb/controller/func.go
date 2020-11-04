package controller

import (
	"MediaWeb/api/dbopts"
	"MediaWeb/api/defs"
	"MediaWeb/api/redisopts"
	"MediaWeb/common"
	"MediaWeb/jwt"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func urlToSliceAndCheck(url string)(dir []string) {
	url = url[1:]
	dir = strings.Split(url, "/")[1:]

	return
}

func dealDbStatus(w *http.ResponseWriter, success string, err error) {
	if err != nil {
		ErrResponse(w, defs.ErrDBOperator(err.Error()))
	} else {
		if err != nil {
			SuccessResponse(w, err.Error())
		}  else {
			SuccessResponse(w, success)
		}

	}
}

// 0 无效，1 有效但需要刷新，2 有效
func checkExp(exp int64) int {
	nowUnix := time.Now().Unix()
	fmt.Println("time:", time.Now(), time.Unix(exp, 0))
	if nowUnix >= exp {
		return 0
	} else {
		if nowUnix - exp < int64(time.Hour) {
			return 1
		}
		return 2
	}
}

func marshalObj(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		panic(err)
		return ""
	} else {
		return string(bs)
	}
}

func refreshToken(u defs.User ,w *http.ResponseWriter) string {
	token, terr := jwt.GenerateVToken(u.Uid, u.Uname)
	if terr != nil {
		panic(terr)
		common.Log.Fatal("hello")
	}
	(*w).Header().Set("X-Token", token)
	return token
}

func clearToken(w *http.ResponseWriter) {
	(*w).Header().Set("X-Token", "")
}

// token,tokenIsValid 0 无效，1 有效，2 有效但需要刷新
func verifyToken(secret string, tokenStr string)(uid int, tokenIsValid int, err error) {
	tokenIsValid = 0

	// token 验证
	if secret != "" {
		var exp int64
		exp, uid, err = jwt.ParseVToken(tokenStr, secret)
		fmt.Println(exp, uid, err)
		if err != nil {
			return
		} else {
			tokenIsValid = checkExp(exp)
		}
	} else {
		err = errors.New("verify failed")
	}

	return
}

func getSecret(uid int)string {
	u := redisopts.GetUserInfo(uid)
	if u == nil {
		var e error
		u, e = dbopts.GetUser(map[string]string{"uid":strconv.Itoa(uid)})
		if e != nil {
			return ""
		} else {
			redisopts.SetUserInfo(*u)
			return u.Uname
		}
	}

	return u.Uname
}


func registerUser(u defs.User)(successStr string ,dbErr error) {
	// user register
	dbErr = dbopts.InsertUser(u)
	if dbErr != nil {
		return
	}

	successStr = marshalObj(u)
	return
}

func loginUser(tokenIsVaild int ,u *defs.User)(successStr string ,dbErr error) {
	if tokenIsVaild == 0 {
		var newU *defs.User
		newU, dbErr = dbopts.GetUser(map[string]string{"uname":u.Uname,"password":u.Password})
		if newU != nil && dbErr == nil {
			redisopts.SetUserInfo(*newU)
			u.Uid = newU.Uid
			successStr = "login success!" + marshalObj(*newU)
		}
	} else {
		successStr = "you are in the login status"
	}

	return
}




