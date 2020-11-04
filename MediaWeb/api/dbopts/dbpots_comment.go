package dbopts

import (
	"MediaWeb/api/defs"
	"MediaWeb/tools/uuid"
	"fmt"
	"time"
)

func InsertComment(info defs.CommentInfo)error {
	info.Cid = uuid.GenerateUUID()
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	sql := "insert into comment values(?,?,?,?,?)"

	err := exec(sql, info.Cid, info.Vid, info.Uid, info.Content, nowtime)

	return err
}

// 只返回评论内容，评论时间，评论者
func ListComment(vid string, from time.Time, to time.Time)(cs []*defs.CommentInfo, err error) {
	sql := `select c.content,c.cdate,u.uname from comment c 
			inner join user u on c.uid = u.uid where
			c.vid = ? and c.cdate > ? and c.cdate <= ?`
	fromFTime := from.Format("2006-01-02 15:04:05")
	ToFTime := to.Format("2006-01-02 15:04:05")

	fmt.Println(fromFTime, ToFTime)
	rs, err := QueryMany(sql, vid, fromFTime, ToFTime)
	if err != nil {
		return
	}
	var c *defs.CommentInfo
	var ctime string
	for i := 0;rs.Next();i++ {
		c = new(defs.CommentInfo)
		rs.Scan(&c.Content, &ctime, &c.Uname)
		c.Cdate, err = time.Parse("2006-01-02 15:04:05", ctime)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			cs = append(cs, c)
		}
	}

	return
}
