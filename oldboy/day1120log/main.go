package main

import (
	"fmt"
	"log"
	"os"
)

/*
	Go语言内置的log包实现了简单的日志服务

	使用logger
		log包定义了logger类型，该类型提供了一些格式化输出的方法
		可以调用Print(Print|Printf|Println）系列、
			   Fatal（Fatal|Fatalf|Fatalln）系列、
			   Panic（Panic|Panicf|Panicln系列


	配置logger
		标准logger的配置
			默认情况下logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，
			比如日志的文件名、行号等，log标准库中提供了定制这些设置的方法
			log标准库中的Flags函数返回标准logger的输出配置，而SetFlags函数用来设置标准
			logger的输出配置
		flag选项
			log标准库提供了如下的flag选项，它们是一系列定义好的常量。
				const (
					// 控制输出日志信息的细节，不能控制输出的顺序和格式。
					// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
					Ldate         = 1 << iota     // 日期：2009/01/23
					Ltime                         // 时间：01:23:23
					Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
					Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
					Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
					LUTC                          // 使用UTC时间
					LstdFlags     = Ldate | Ltime // 标准logger的初始值
				)

		配置日志前缀
			log标准库中还提供了关于日志信息前缀的两个方法
				func Prefix() string//查看前缀
				func SetPrefix(prefix string)//设置前缀

		配置日志输出位置
			SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出

		创建logger
			log标准库中还提供了一个创建新logger对象的构造函数-New，支持我们创建自己的logger实例
				func New(out io.Writer,prefix string,flag int) *Logger
				logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)


	Go内置的log库功能有限，无法满足记录不同级别日志的情况，实际项目中一般采用第三方
	的日志库，如logrus、zap等
*/
func init() {
	logFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Ltime)
}
func main() {
	log.Println("this is a normal log")
	log.SetPrefix("hello")
	// log.Fatalln("this is a fatal log")//Fatal系列函数会在后边调用os.Exit(1)
	// log.Panicln("this is a panic log")//Panic系列函数会在写入日志信息后panic
	log.Println("this is a normal log2")
}
