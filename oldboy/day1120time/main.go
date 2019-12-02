package main

import (
	"fmt"
	"time"
)

/*
	time包提供了时间的显示和测量用的函数，日历采用公历

	time.Time类型表示时间，我们可以通过time.Now()函数获取当前的时间对象，然后获取时间
	对象的年月日时分秒等信息

	时间戳：自1970年1月1日至当前时间的总毫秒数，也被称为Unix时间戳

	时间间隔
	time.Duration是time包定义的一个类型，它表示两个时间点之间经过的时间，以纳秒为单位
	表示一段时间间隔，最长大约是290年
	time.Duration表示1纳秒

	时间操作
		Add:时间+时间间隔，多久之后
		Sub:两个时间之间的差值
		Equal:判断两个时间是否相同
		Before
		After
		定时器:time.Tick(间隔),返回一个channel
		时间格式化:now.Format("2006/01/02 15:04")

*/
func main() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	t1 := now.Add(1)
	t1u := t1.UnixNano()
	fmt.Println(timestamp2)
	fmt.Println(t1u)
	d := t1.Sub(now)
	fmt.Println(d)
	fmt.Println(now.Equal(t1))
	tickDemo()
}
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}
