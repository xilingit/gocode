package main

import (
	"encoding/json"
	"fmt"
)

/*
嵌套结构体
嵌套匿名结构体
	要避免嵌套结构体的字段名冲突

结构体的“继承”
	Go语言中使用结构体也可以实现其他语言中面向对象的继承

结构体字段的可见性
	结构体中字段大写开头表示可公开访问，小写表示私有(仅在定义当前结构体的包中可访问)

结构体与JSON序列化
	JSON:是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成

结构体标签Tag
	Tag是结构体的元信息，可以在运行的时候通过反射机制读取出来。
	Tag在结构体字段的后方定义，由一对反引号包裹起来，具体格式如下:
		`key1:"value1" key2:"value2"`
	规则:结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔
*/
//Animal struct
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

//Dog struct
type Dog struct {
	age int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s它会汪汪叫\n", d.name)
}

func main() {
	d := Dog{
		age: 10,
		Animal: &Animal{
			name: "小煌",
		},
	}
	d.move()
	d.wang()

	c := &Class{
		Title:    "biaotou",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 2; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "nan",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

//Student struct
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

//Class 结构体
type Class struct {
	Title    string     `json:"title"`
	Students []*Student `json:"students"`
}
