package main

import "fmt"

type person struct {
	name string  //姓名
	idType string  //身份证类型
	id string    //身份证
}

//get和set方法
func (p *person)getName() string {
	return p.name;
}

func (p *person)getIdType() string {
	return p.idType
}

func (p *person)getId() string {
	return p.id
}

func (p *person)setAll(data map[string]string) {
	for key, value := range data {
		if key == "name" {
			p.name = value
		} else if key == "idType" {
			p.idType = value
		} else if key == "id" {
			p.id = value
		}
	}
}

//打印信息
func (p *person)printPerson() {
	fmt.Println("name：", p.name)
	fmt.Println("idType：", p.idType)
	fmt.Println("id：", p.id)
}
