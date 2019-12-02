package main

import (
	"fmt"
	"mytestgo/oldboy/day1125protobuf/address"
)

/*
	protobuf是一种高效的数据格式，平台无关、语言无关、可扩展，可用于RPC系统和持续数据存储系统

	protobuf是一个与语言无关的一个数据协议，所以我们需要先编写IDL文件然后借助工具生成指定语言的代码，
	从而实现数据的序列化与反序列化过程
	大致开发流程如下：1、编写IDL 	2、生成指定语言的代码		3、序列化和反序列化

	编译器安装
*/
func main() {
	// var cb address.ContactBook

	p1 := address.Person{
		Name:   "小王子",
		Gender: address.GenderType_MALE,
		Number: "7878778",
	}
	fmt.Println(p1)
	// cb.Persons = append(cb.Persons, &p1)
	// // 序列化
	// data, err := proto.Marshal(&p1)
	// if err != nil {
	// 	fmt.Printf("marshal failed,err:%v\n", err)
	// 	return
	// }
	// ioutil.WriteFile("./proto.dat", data, 0644)

	// data2, err := ioutil.ReadFile("./proto.dat")
	// if err != nil {
	// 	fmt.Printf("read file failed, err:%v\n", err)
	// 	return
	// }
	// var p2 address.Person
	// proto.Unmarshal(data2, &p2)
	// fmt.Println(p2)
}
