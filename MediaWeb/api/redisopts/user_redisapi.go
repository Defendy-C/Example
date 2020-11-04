package redisopts

import (
	"MediaWeb/api"
	"MediaWeb/api/defs"
	"MediaWeb/common"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
)

func exec(comm ,key string, args interface{})(error) {
	var err error
	as := redis.Args{}
	as = as.Add(key)
	if strings.ToLower(comm[len(comm) - 2:]) == "ex" {
		as = as.Add(common.TimeOut.Seconds())
	}

	_, err = api.RedisWriteConn.Conn.Do(comm, as.AddFlat(args)...)
	return err
}

func query(comm, key string, fields ...interface{}) (res interface{}, err error) {
	var args = redis.Args{}
	args = args.Add(key)
	for _, v := range fields {
		args = args.AddFlat(v)
	}
	res, err = api.RedisReadConn.Conn.Do(comm, args...)

	return
}

func GetUserInfo(uid int)*defs.User {
	var u defs.User

	var res []byte
	var err error
	var comm string = "get"
	res, err = redis.Bytes(query(comm, strconv.Itoa(uid)))
	if err != nil {
		return nil
	}

	err = json.Unmarshal(res, &u)
	if err != nil {
		return nil
	}

	return &u
}

func SetUserInfo(user defs.User) {
	bs, err := json.Marshal(user)
	if err != nil {
		common.Log.Error(err.Error())
		return
	}

	str := string(bs)
	err = exec("setex",strconv.Itoa(user.Uid), str)
	if err != nil {
		common.Log.Error(err.Error())
	}
}

