package main

type loginItem struct {
	 yhm string   //用户名
	 mm string    //密码
	 yzm string   //验证码
}

func (li *loginItem) setYHM(yhm string) {
	li.yhm = yhm
}

func (li *loginItem) setMM(mm string) {
	li.mm = mm
}

func (li *loginItem) setYZM(yzm string) {
	li.yzm = yzm
}

func (li *loginItem) getYHM() string {
	return li.yhm
}

func (li *loginItem) getMM() string {
	return li.mm
}

func (li *loginItem) getYZM() string {
	return li.yzm
}
