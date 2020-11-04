package main

import "fmt"

type courseIm struct {
	name string   //课程名
	teacher string   //老师
	credit string   //学分
	score string    //分数
	point string    //绩点
}

//打印信息
func (im *courseIm)printCourse() {
	fmt.Printf("%v  %v  %v  %v  %v\n",im.name, im.teacher,im.credit,im.score,im.point)
}

//获取信息，根据key值存储信息
func (im *courseIm)setIm(item interface{}, key string) bool {
	if key == "kcmc" {
		im.name = item.(string)
		return true
	} else if key == "xf" {
		im.credit = item.(string)
		return true
	} else if key == "cj" {
		im.score = item.(string)
		return true
	} else if key == "jd" {
		im.point = item.(string)
		return true
	} else if key == "jsxm" {
		im.teacher = item.(string)
		return true
	}
	return false
}
