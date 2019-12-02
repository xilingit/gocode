package main

import (
	"fmt"
	"strings"
)

/*
Go语言提供的映射关系容器为map，内部使用散列表hash实现
map是一种无序的基于key-value的数据结构，是引用类型，必须初始化才能使用
定义: map[key类型] value类型
使用make()函数来分配内存，语法为
		make(map[key类型]value类型，[容量])

判断map中key是否存在的特殊写法，格式如下：
	value,ok := map[key]

map遍历
	使用for range
	for k,v := range map 	//k,v一起遍历
	for k := range map	 	//只遍历k

删除map键值对
	使用内置的delete()函数删除键值对
		delete(map,key)
*/
func main() {
	// scoreMap := make(map[string]int, 3)
	// scoreMap["语文"] = 96
	// scoreMap["数学"] = 97
	// scoreMap["英语"] = 98
	// scoreMap["化学"] = 99
	// fmt.Println(scoreMap) //无序的

	// userInfo := map[string]string{
	// 	"username": "一只小喵",
	// 	"password": "123456",
	// }
	// fmt.Println(userInfo) //声明的时候填充元素

	// if v, ok := scoreMap["化学"]; ok {
	// 	fmt.Printf("key=化学,value=%d\n", v)
	// } else {
	// 	fmt.Println("map中不存在化学分数")
	// }
	// delete(scoreMap, "英语")
	// fmt.Println(scoreMap)

	// rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	// var scoreMap = make(map[string]int, 100)

	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	// 	value := rand.Intn(100)          //生成0~99的随机整数
	// 	scoreMap[key] = value
	// }
	// //取出map中的所有key存入切片keys
	// var keys = make([]string, 0, 100)
	// for key := range scoreMap {
	// 	keys = append(keys, key)
	// }
	// //对切片进行排序
	// sort.Strings(keys)
	// //按照排序后的key遍历map
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap[key])
	// }

	// var mapSlice = make([]map[string]string, 3)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
	// fmt.Println("after init")
	// // 对切片中的map元素进行初始化
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "小王子"
	// mapSlice[0]["password"] = "123456"
	// mapSlice[0]["address"] = "沙河"
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }

	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	var str = "how do you do"
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)
	var strMap map[string]int
	strMap = make(map[string]int, 10)
	fmt.Println(strMap)

}
