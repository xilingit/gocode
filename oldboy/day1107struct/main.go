package main

import (
	"fmt"
)

/*
	Go语言中没有类的概念，也不支持类的继承等面向对象的概念，Go语言中通过
	结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性

	自定义类型
		type MyInt int //将MyInt定义为int类型
		通过type关键字的定义，MyInt是一种新的类型，具有int的特性

	类型别名
		是Go1.9版本添加的新功能
		type TypeAlias = Type //TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型
		比如系统的
				type byte = int8
				type rune = int32


	结构体
		封装多个基本数据类型，struct来定义自己的类型，进而实现面向对象
		定义格式
			type 类型名 struct{
				字段名 字段类型
				......
			}
		结构体实例化时才会真正地分配内存，声明结构体类型格式
			var 结构体实例  结构体类型

		匿名结构体
			在定义一些临时数据结构等场景下可以使用匿名结构体
			var user struct{Name string; Age int}
			user.Name = "xyz"
			user.Age = 22
			fmt.Println(user)

		创建指针类型结构体
			我们可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址，格式如下：
				var p2 = new(Person)
				fmt.Printf("%T\n",p2) //*main.Person
				fmt.Printf("p2=%#v\n",p2) //p2=&main.person{name:"", city:"", age:0}
			Go语言中支持对结构体指针直接使用.来访问结构体的成员
			使用&对结构体进行取地址相当于对结构体类型进行了一次new实例化操作

		结构体初始化
			没有初始化的结构体，其成员变量都是的对应的零值

		使用键值对初始化
			使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
			也可以对结构体指针进行键值对初始化
			当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
		使用值的列表初始化
			初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
				注意
						必须初始化结构体的所有字段。
						初始值的填充顺序必须与字段在结构体中的声明顺序一致。
						该方式不能和键值初始化方式混用。

		结构体内存布局
			结构体占用一块连续的内存


		构造函数
			Go语言的结构体没有构造函数，我们可以自己实现。如下：
				func newPerson(name, city string, age int8) *person {
					return &person{
						name: name,
						city: city,
						age:  age,
					}//直接返回指针，性能开销小
				}

		方法和接收者
			Go语言中的方法是一种作用于特定类型变量的函数，这种特定类型变量叫做接收者，接收者
				的概念类似于其他语言中的this或self
			格式如下:
				func (接收者变量 接收者类型) 方法名(参数列表) (返回参数){
					函数体
				}
			方法和函数的区别是，函数不属于任何类型，方法属于特定类型
			值类型接收者使用情况
				1、需要修改接收者中的值
				2、接收者是拷贝代价比较大的大对象
				3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

			任意类型添加方法
				方法的接收者可以是任意类型，不仅仅是结构体，任何类型都可以拥有方法，比如下边的MyInt,NewInt不行
				非本地类型不能定义方法，我们不能给别的包的类型定义方法


		结构体匿名字段
			结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段
			//Person 结构体Person类型
			type Person struct {
				string
				int
			}

			p1 := Person{
				"小王子",
				18,
			}
			fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
			fmt.Println(p1.string, p1.int) //北京 18

*/

//自定义类型
type MyInt int

//类型别名
type NewInt = int

//Person的自定义类型
type Person struct {
	age  int8
	city string
	name string
}

func newPerson(name string, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

//Person.Dream()的方法
func (p *Person) Dream() {
	fmt.Printf("%s的梦想是干%s\n", p.name, p.city)
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	var a MyInt
	var b NewInt
	fmt.Printf("type of a:%T\n", a) //main.MyInt
	fmt.Printf("type of b:%T\n", b) //int

	var p Person
	p.name = "guan"
	p.city = "beijing"
	p.age = 33
	fmt.Println(p)

	var user struct { //匿名结构体
		Age  int
		Name string
	}
	user.Name = "xyz"
	user.Age = 22
	fmt.Println(user)

	var p2 = new(Person)
	fmt.Printf("type of p2 is :%T\n", p2) //*main.Person
	p2.city = "aa"
	p2.name = "hello"
	p2.age = 22
	fmt.Println(p2) //&{aa hello 22}
	p3 := newPerson("大黄", "河南", 23)
	p3.Dream()

	var t test
	t = test{
		1, 2, 3, 4,
	}
	fmt.Printf("t.a %p\n", &t.a)
	fmt.Printf("t.b %p\n", &t.b)
	fmt.Printf("t.c %p\n", &t.c)
	fmt.Printf("t.d %p\n", &t.d)

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		// stu1 := stu  //打开这个后就是一一对应了
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		// 娜扎 => 大王八
		// 大王八 => 大王八
		// 小王子 => 大王八
	}

	//匿名字段默认采用类型名作为字段名
	//结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	niMing := NiMing{
		string: "sss",
		int:    33,
	}
	fmt.Println(niMing)
}

type student struct {
	age  int
	name string
}
type NiMing struct {
	int
	string
}
