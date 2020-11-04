package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var cookies []*http.Cookie
var cookiesJar http.CookieJar

//创建client
var client *http.Client

//用于存储url
var URL map[string]interface{}

//token.cookies
var t token = token{
	csrftoken: "",
}

//公钥和获取公钥的cookies
var key PublicKey = PublicKey{
	modulus:  "",
	exponent: "",
}

//用户名，密码，验证码，验证码用不到
var li loginItem = loginItem{
	yhm: "",
	mm:  "",
	yzm: "",
}

//个人信息：姓名，身份证类型，身份证号
var p person = person{
	name:   "",
	idType: "",
	id:     "",
}

//初始化
func init() {
	URL = make(map[string]interface{})
	cookies = nil
	cookiesJar, _ = cookiejar.New(nil)
	client = &http.Client{Jar:cookiesJar}
}

//检查出错函数
func checkError(err error) {
	if(err != nil) {
		log.Fatal(err)
	}
}

//生成目前时间戳
func getNowTimeStamp() string {
	return strconv.FormatInt((time.Now().Unix()), 10)
}

//打印响应信息的函数，用于测试
func showResponseImformation(name string, data []byte, resp *http.Response) {
	fmt.Println(name+":")
	fmt.Println("header:", resp.Header)
	fmt.Println("status:", resp.Status)
	fmt.Println("request:", resp.Request)
	//fmt.Println(resp.Request.PostForm)

	fh,err := os.Create(name)
	defer fh.Close()
	checkError(err)
	fh.WriteString(string(data))
}

func directjm(password string)(mm string) {
	//获取公钥
	key.setPublicKey(URL["getPublicKey"].(string))
	//key.printPublicKey()
	//fmt.Println()

	//密码加密及用户名获取
	nstr := b64tohex(key.modulus)
	estr := b64tohex(key.exponent)
	var rsajm  RSAPublic
	mm = hex2b64(rsajm.RSAEncrypt(password, nstr, estr))
	return
}

func post() {
	//封装postData
	postValue := url.Values{"csrftoken":{t.getToken()}, "yhm":{li.getYHM()}, "mm":{li.getMM()}}
	postString := postValue.Encode()

	//设置请求
	var httpReq *http.Request
	httpReq, err := http.NewRequest("POST", URL["login"].(string), bytes.NewBufferString(postString))
	checkError(err)

	//设置请求头
	httpReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0")
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=uft-8")

	//执行请求
	_, err = client.Do(httpReq)
	checkError(err)
	cookies = cookiesJar.Cookies(httpReq.URL)

	//打印响应数据
	//data, err := ioutil.ReadAll(resp.Body)
	//checkError(err)
	//fmt.Println()
	//showResponseImformation("post.html", data, resp)
}

//获取个人信息
func getPersonIm() {
	//map类型用于存储个人信息
	var pMap map[string]string
	pMap = make(map[string]string)

	//从网络上获取html
	request, err := http.NewRequest("GET", URL["person"].(string), nil)
	checkError(err)
	resp, err := client.Do(request)
	checkError(err)
	html, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	//正则表达式提取数据
	re, err := regexp.Compile(`<p class="form-control-static">([^a-zA-Z]*)?</p>`)
	checkError(err)
	data := re.FindAllStringSubmatch(string(html), -1)

	/*
	index索引
	其中1为学号，2，3为姓名，4为身份证类型，5为身份证
	 */
	index := 1
	for _, value := range data {
		if index == 3 {
			index++
			continue
		}
		if index == 6 {
			break
		}
		str := strings.Trim(value[1], "\r\n")
		str = strings.Trim(str, "\t")
		str = strings.Trim(str, "\r\n")
		if str != "" {
			if index == 2 {
				pMap["name"] = str
			} else if index == 4 {
				pMap["idType"] = str
			} else if index == 5 {
				pMap["id"] = str
			}
			index++
		}
	}

	//存储个人信息数据
	p.setAll(pMap)
}

//从接口数组中获取课程信息
func fillCourseData(im interface{}, container []courseIm, tag int) {
	data := im.([]interface{})
	for key, value :=range data {
		valueMap := value.(map[string]interface{})
		for k, v := range valueMap {
			container[key].setIm(v, k)
		}
	}
}

func getAndPrintCourse() {
	//从网络上提取json文件
	postValue := url.Values{"xh_id":{li.yhm},"xnm":{""},"xqm":{""},"_search":{"false"},"nd":{getNowTimeStamp()},"queryModel.showCount":{"50"},"time":{"0"}}
	postString := postValue.Encode()
	request, err := http.NewRequest("POST", URL["course"].(string), bytes.NewBufferString(postString))
	checkError(err)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp,err := client.Do(request)
	data, err := ioutil.ReadAll(resp.Body)

	//将json格式转换为接口格式
	var courseData interface{}
	json.Unmarshal(data, &courseData)
	//fmt.Println()
	//showResponseImformation("course.json", data, resp)

	//获取课程总数
	var size int
	courseDataMap := courseData.(map[string]interface{})
	for key,value := range courseDataMap {
		if key == "totalResult" {
			size = int(value.(float64))
			break
		}
	}

	//定义courseIm数组用于存储成绩信息
	courseItem := make([]courseIm, size)
	//填充数组
	fillCourseData(courseDataMap["items"], courseItem, 0)
	//打印数组
	fmt.Println()
	fmt.Println("name  teacher  credit  score  point")
	for i:=0;i<size;i++ {
		courseItem[i].printCourse()
	}
}

func main() {
	//域名
	var Url string = "http://jwcjwxt.jyu.edu.cn"

	URL["token"] = Url + "/xtgl/login_slogin.html?language=zh_CN&_t="+getNowTimeStamp()
	URL["getPublicKey"] = Url + "/xtgl/login_getPublicKey.html?time=" + getNowTimeStamp()
	URL["login"] = Url + "/xtgl/login_slogin.html"
	URL["course"] = Url + "/cjcx/cjcx_cxDgXscj.html?doType=query&gnmkdm=&su="
	URL["person"] = Url + "/xsxxxggl/xsgrxxwh_cxXsgrxx.html?gnmkdm=N100801&layout=default&su="

	//生成token,并获取cookies
	t.initTokenAndCookies(URL["token"].(string))
	//t.printToken()
	//fmt.Println()

	//获取用户名和密码
	var user string
	var password string
	fmt.Printf("请输入用户名：")
	fmt.Scan(&user)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Println()

	//直接加密并赋值
	li.setYHM(user)
	li.setMM(directjm(password))

	//提交数据
	post()

	//获取个人信息
	getPersonIm()

	//打印个人信息
	p.printPerson()

	//获取课程信息并打印
	getAndPrintCourse()
}