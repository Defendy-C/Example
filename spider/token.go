package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type token struct {
	csrftoken string
}

//初始化token和cookies
func (t *token) initTokenAndCookies(url string) {
	//设置请求
	request, err := http.NewRequest("GET", url, nil)
	checkError(err)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0")
	resp, err := client.Do(request)
	defer resp.Body.Close()
	checkError(err)

	//获取登录页以用正则表达式提取token
	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	//showResponseImformation("getCsrftoken.html", data, resp)

	zz ,err := regexp.Compile(`<input type="hidden" id="csrftoken" name="csrftoken" value="(.+)"/>`)
	checkError(err)
	findArray := zz.FindStringSubmatch(string(data))
	t.csrftoken = findArray[1]

	//获取cookies
	cookiesJar.Cookies(request.URL)
}

//打印token
func (tac *token) printToken() {
	fmt.Println("csrftoken:", tac.csrftoken, "size:", len(tac.csrftoken))
}

//获取token
func (t *token) getToken() string {
	return t.csrftoken
}
