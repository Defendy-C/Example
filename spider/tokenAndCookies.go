package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
)

type tokenAndCookies struct {
	csrftoken string
	cookies []*http.Cookie
	jar *cookiejar.Jar
	client *http.Client
}

func (tac *tokenAndCookies) initTokenAndCookies(url string) {
	//jar 和 client
	tac.jar, _ = cookiejar.New(nil)
	tac.client = &http.Client{Jar:tac.jar}

	request, err := http.NewRequest("GET", url, nil)
	checkError(err)
	//request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0")
	resp, err := tac.client.Do(request)
	defer resp.Body.Close()
	checkError(err)

	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Println("url:", url)
	showResponseImformation("getCsrftoken.html", data, resp)

	zz ,err := regexp.Compile(`<input type="hidden" id="csrftoken" name="csrftoken" value="(.+)"/>`)
	checkError(err)
	findArray := zz.FindStringSubmatch(string(data))
	tac.csrftoken = findArray[1]
	tac.cookies = tac.jar.Cookies(request.URL)
}

func (tac *tokenAndCookies) printToken() {
	fmt.Println("csrftoken:", tac.csrftoken, "size:", len(tac.csrftoken))
}

func (tac *tokenAndCookies) printCookies() {
	fmt.Println("cookies:")
	for _, value := range tac.cookies {
		fmt.Println(value)
	}
}

func (tac *tokenAndCookies) getToken() string {
	return tac.csrftoken
}

func (tac *tokenAndCookies) getCookies() []*http.Cookie {
	return tac.cookies
}

func (tac *tokenAndCookies) setCookies(cookies [] *http.Cookie) {
	tac.cookies = cookies
}