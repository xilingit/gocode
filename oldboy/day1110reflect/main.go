package main

import (
	"fmt"
	"reflect"
)

/*
	反射:指在程序运行期对程序本身进行访问和修改的能力

	reflect包
		任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且提供了
		reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type

	TypeOf
		使用reflect.TypeOf()函数可以获得任意值的类型对象

	在反射中关于类型划分为两种，类型(Type)和种类(Kind),种类就是指底层类型
	type myInt int64 //type:myInt	kind:int64
	如果是结构体，type:结构体名		kind:struct
	Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空

	ValueOf
		使用ValueOf()返回的是reflect.Value类型，包含原始值的信息


	isNil()和isValid()
		func (v Value) IsNil() bool
		报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一，否则导致panic
		常被用于判断指针是否为空

		func (v Value) IsValid() bool
		返回v是否持有一个值，如果v是Value零值会返回假
		常被用于判定返回值是否有效

	与结构体相关的方法
		反射对象如果是结构体，可以通过NumField()和Field()方法获得结构体成员的详细信息


	反射是把双刃剑
	反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
	1、基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
	2、大量使用反射的代码通常难以理解。
	3、反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。

*/
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	v.Kind()
	v.Name()
	fmt.Printf("type of v is %v\n", v)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func main() {
	var aa float32 = 3.14
	reflectType(aa) //type of v is float32
	var bb int64 = 100
	reflectType(bb) //type of v is int64

	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value

	//通过反射设置变量的值
	var d int64 = 100
	// reflectSetValue(&d) //不会修改d的值
	reflectSetValue2(&d)
	fmt.Println(d)

	//IsNil()和IsValid()
	var p *int
	fmt.Println("var p *int IsNil:", reflect.ValueOf(p).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	p1 := struct {
		abcd string
	}{abcd: "hello"}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(p1).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(p1).MethodByName("abc").IsValid())
	// map
	p2 := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(p2).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	var ee int8 = 44
	fmt.Println("ee type of ", reflect.TypeOf(ee))
	fmt.Println("ee value of ", reflect.ValueOf(ee))

	stu1 := student{
		Name:  "xiao guan",
		Score: 33,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	//反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
