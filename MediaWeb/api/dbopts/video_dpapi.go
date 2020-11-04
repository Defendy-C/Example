package dbopts

import (
	"MediaWeb/api/defs"
	"MediaWeb/tools/uuid"
	"time"
)

// v必须有id和name
func InsertVideoInfo(v defs.VideoInfo)error {
	v.Vid = uuid.GenerateUUID()
	vtimestamp := time.Now().Format("2006-01-02 15:04:05")
	sql := "insert into video values(?, ?, ?, ?)"
	err := exec(sql, v.Vid, v.Uid, v.Vname, vtimestamp)

	return err
}

func GetVideoInfo(uid int) (v *defs.VideoInfo, err error) {
	v = nil
	err = nil

	sql := "select * from video where uid=?"
	r, err := QueryOne(sql, uid)
	if err != nil {
		return
	}

	v = new(defs.VideoInfo)
	var timestr string
	r.Scan(&v.Vid, &v.Uid, &v.Vname, &timestr)
	v.Vdate, err = time.Parse("2006-01-02 15:04:05", timestr)

	return
}

func DelVideo(vid string) error {
	sql := "delete from video where vid=?"
	err := exec(sql , vid)

	return err
}
