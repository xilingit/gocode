package main

import "net/http"

import "fmt"

import "io/ioutil"

import "net/url"

/*
	Go语言内置的net/http包十分优秀，提供了HTTP客户端和服务端的实现

*/
func main() {

	//一个简单的发送http请求的client端
	resp, err := http.Get("http://www.liwenzhou.com/")
	if err != nil {
		fmt.Println("get request error,error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error,error:", err)
		return
	}
	fmt.Println(string(body))

	apiUrl := "http://127.0.0.1:9090/get"
	fmt.Println(apiUrl)
	data := url.Values{}
	data.Set("name","ss")
	data.Set("age","18")
	url.ParseRequestURI(apiUrl)
}
