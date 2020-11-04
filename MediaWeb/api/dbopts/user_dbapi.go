package dbopts

import (
	"MediaWeb/api"
	"MediaWeb/api/defs"
)

// u 必须有username和password
func InsertUser(u defs.User) error {
	sql := "insert into user values(null, ?, ?)"
	err := exec(sql, u.Uname, u.Password)

	return err
}

func GetUser(where map[string]string) (u *defs.User, err error) {
	u = new(defs.User)
	sql := "select uid,uname from user where"
	var args []interface{}
	i := 0
	for k, v := range where {
		sql += " " + k + "=? and"
		args = append(args, v)
		i++
	}
	sql = sql[:len(sql) - 3]
	r, err := QueryOne(sql, args...)
	if err != nil {
		return
	}

	r.Scan(&u.Uid, &u.Uname)
	if err == api.QueryNilErr {
		err = nil
	}

	return
}

func DelUser(username string) error {
	sql := "delete from user where uname=?"
	err := exec(sql, username)


	return err
}

func UpdateUser() bool {
	// TODO
	return false
}





