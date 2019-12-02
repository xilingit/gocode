package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	Go语言中的模版
	html/template包实现了数据驱动的模板，用于生成可对抗代码注入的安全html输出，它提供了和
	text/template包相同的接口，Go语言中输出html的场景都应使用text/template包

	模板
		模板语法
			都包含在{{和}}中间，其中{{.}}中的点表示当前对象
			当我们传入一个结构体对象时，我们可以根据.来访问结构体的对应字段
*/
type UserInfo struct {
	Name, Gender string
	Age          int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	userInfo := UserInfo{
		Name :"王子",
		Gender: "男",
		Age:23,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tpl.Execute(w, userInfo)
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
