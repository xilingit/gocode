package main

import "fmt"

type Shaper interface {
	//定义接口
	Areas() float32
}

type Square struct {
	//结构体
	side float32
}

func (sq *Square) Areas() float32 { //接口方法
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square) //new一个结构体对象
	sq1.side = 5
	var areaIntf Shaper //接口
	fmt.Printf("areaIntf的类型是%T\n", areaIntf)
	areaIntf = sq1 //方式一
	fmt.Printf("sq1的类型是%T,   areaIntf的类型是%T\n", sq1, areaIntf)
	//areaIntf = Shaper(sq1) //方式二
	fmt.Printf("The square has area: %f\n", areaIntf.Areas())

	bird := new(Bird)
	bird.color = "黄色"
	var flyInterface Fly
	flyInterface = bird
	fmt.Println(flyInterface.Fly())
	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
	fmt.Println(1024 * 1024)
	fmt.Println(Name)
}

type Bird struct {
	color string
}

func (bird *Bird) Bark() int {
	panic("implement me")
}

type Fly interface {
	Fly() string
	Bark() int
}

func (bird *Bird) Fly() string {
	fmt.Println(bird.color, "郑子啊飞")
	return bird.color + "fei"
}
