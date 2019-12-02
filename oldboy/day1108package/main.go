package main

import (
	"fmt"
	cal "mytestgo/oldboy/day1108package/add" //引用自己的包，别名可写可不写

	_ "github.com/beego/bee/utils" //引用第三方包
)

/*
	包是多个Go源码的集合，主要用于代码复用，比如内置包fmt、os、io等
	包可以理解为存放.go文件的文件夹

	注意:
		一个文件夹下面只能有一个包，同样一个包的文件不能在多个文件夹下。
		包名可以不和文件夹的名字一样，包名不能包含-符号。
		包名为main的包为应用程序的入口包，编译时不包含main包的源代码时不会得到可执行文件。

	可见性
		将标识符的首字母大写对外可见，小写对外不可见

	包的导入
		单行导入  	import "包名"
		多行导入	import ("包名")
		匿名导入	import _ "包的路径"

	init()初始化
		Go语言导入包语句会自动触发包内部的init()函数
		init()函数没有返回值，没有参数，不能在代码中主动调用它
		全局声明---->init()---->main()
*/
func main() {
	fmt.Println(cal.A)
	fmt.Println(cal.Sum(3, 4))
}
