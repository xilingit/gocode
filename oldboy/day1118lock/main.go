package main

import (
	"fmt"
	"sync"
	"time"
)

/*

	并发安全和锁
		多个goroutine同时操作一个资源，会发生竞态问题

	互斥锁
		互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
		Go语言中使用sync包的Mutex类型来实现互斥锁
		使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
		当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的

	读写互斥锁
		互斥锁是完全互斥的，但是很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候
		是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型

		读写锁分为两种，读锁和写锁。
			当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取
			写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待
			非常适合读多写少的场景

	sync.WaitGroup
		在代码中生硬的使用time.Sleep肯定不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步
			wg.Add(delta int)//计数器+delta
			wg.Done()//计数器-1
			wg.Wait()//阻塞直到计数器变为0
		WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
		例如当我们启动了N个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将
		计数器减1.通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成
		WaitGroup是个结构体，传递的时候要传递指针


	sync.Once
		当我们需要确保某些操作在高并发的场景下只执行一次，例如只加在一次配置文件、只关闭一次通道等
		Go语言中的sync包中提供了一个针对只执行一次场景的解决方案--sync.Once
		只有一个Do方法
			func (o *Once) Do(f func()){}
*/
// func add() {
// 	for i := 0; i < 500; i++ {
// 		lock.Lock() //使用互斥锁
// 		x = x + 1
// 		lock.Unlock() //使用互斥锁
// 	}
// 	wg.Done()
// }

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
