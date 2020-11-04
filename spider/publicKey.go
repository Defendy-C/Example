package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PublicKey struct {
	modulus string
	exponent string
}

//用于获取公钥
func (key *PublicKey)setPublicKey(url string) {
	//获取存储公钥的json页面
	request, err := http.NewRequest("GET", url, nil)
	checkError(err)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(request)
	checkError(err)
	defer resp.Body.Close()

	//获取的文件为json文件
	data, err := ioutil.ReadAll(resp.Body)
	//showResponseImformation("getPublicKey.json",data, resp)

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	checkError(err)
	//存储公钥
	jsonDataMap := jsonData.(map[string]interface{})
	key.modulus = jsonDataMap["modulus"].(string)
	key.exponent = jsonDataMap["exponent"].(string)
}

//公钥get方法
func (key *PublicKey)getPublicKey() (modulus string, exponent string) {
	if key.modulus == "" && key.exponent == "" {
		fmt.Println("error : undefined")
		return "", ""
	}
	modulus = key.modulus
	exponent = key.exponent
	return
}

func (key *PublicKey)printPublicKey() {
	fmt.Println("modulus:", key.modulus, "size:", len(key.modulus))
	fmt.Println("exponent:", key.exponent, "size", len(key.exponent))
}


