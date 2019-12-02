package day1119test

import "strings"

/*
	Go语言中的测试依赖go test命令，编写测试代码和编写普通的Go代码过程是类似的，不需要学习新的语法、规则或工具
	在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中去

	在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数

		类型		格式					  作用
		测试函数	函数名前缀为Test			测试程序的一些逻辑行为是否正确
		基准函数	函数名前缀为Benchmark		测试函数的性能
		示例函数	函数名前缀为Example			为文档提供示例文档

	测试函数：必须导入testing包
		格式 func TestName(t *testing.T){}
		测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头,其中参数t用于报告测试失败和附加的日志信息

	基准函数
		在一定的工作负载之下检测程序性能的一种方法，基本格式如下
			func BenchmarkName(b *testing.B){
				//
			}
		以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试
		才有对照性	

	示例函数
		它们的函数是以Example为前缀，它们既没有参数也没有返回值
		格式如下
			func ExampleName(){}

*/
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
