package main

import "fmt"

/*
	并发与并行
		并发：同一时间段内执行多个任务
		并行：同一时刻执行多个任务

	goroutine:类似线程，属于用户态线程，由Go语言的运行时调度完成，而线程是由操作系统调度完成
	channel在多个goroutine间进行通信，goroutine和channel是Go语言秉承的CSP并发模式的重要实现基础
	Go语言在语言层面上已经内置了调度和上下文切换的机制

	为一个函数创建一个goroutine，函数前加上go就可以


	Go语言中的操作系统线程和goroutine的关系：
		一个操作系统线程对应用户态多个goroutine。
		go程序可以同时使用多个操作系统线程。
		goroutine和OS线程是多对多的关系，即m:n。

	channel
		CSP模型提倡通过通信共享内存
		通过这个channel实现goroutine之间的通信
		声明格式如下:
			var 变量 chan 元素类型
		创建channel
			引用类型，空值是nil
			声明后需要使用make函数初始化之后才能使用
			var ch chan int = make(chan int,5)//缓冲大小可写可不写
		操作channel
			发送、接收、关闭
			发送和接收都用<-符号
				发送   ch <- 10 //将10发送到ch中
				接收   x := <- ch //从ch中接收值并赋值给变量x
					  <- ch		//从ch中接收值，忽略结果
				关闭   close(ch)

	无缓冲通道
		在make时不指定通道的容量，发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，
		两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个
		goroutine在该通道上发送一个值，导致发送和接收的goroutine同步化，也称同步通道

	有缓冲的通道
		在make时指定通道的容量

	单向通道
		需要对通道进行限制，只能发送或接收
*/

// var wg sync.WaitGroup

// func hello(i int) {
// 	defer wg.Done() //goroutine结束就登记-1
// 	fmt.Println("hello goroutine", i)
// }

// func recv(c chan int) {
// 	ret := <-c
// 	fmt.Println(ret)
// }

func main() {
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) //启动一个goroutine就登记+1
	// 	go hello(i)
	// }
	// wg.Wait() //等待所有登记的goroutine都结束

	// ch := make(chan int)
	// go recv(ch)
	// ch <- 10
	// fmt.Println("发送成功")

	ch1 := make(chan int)
	ch2 := make(chan int)
	// //开启goroutine将0～100的数发送到ch1中
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		ch1 <- i
	// 	}
	// 	close(ch1)
	// }()
	// //开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	// go func() {
	// 	for {
	// 		i, ok := <-ch1
	// 		if !ok {
	// 			break
	// 		}
	// 		ch2 <- i * i
	// 	}
	// 	close(ch2)
	// }()
	// //在主goroutine中从ch2中接收值打印
	// for i := range ch2 {
	// 	fmt.Println(i)
	// }

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func counter(in chan<- int) { //chan<- 只能发送的通道，不能接收
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
}

func squarer(in chan<- int, out <-chan int) { //<-chan 只能接收的通道，不能发送
	for i := range out {
		in <- i * i
	}
	close(in)
}

func printer(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
