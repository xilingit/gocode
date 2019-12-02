package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	sync.Map
		Go语言中内置的map不是并发安全的
		Go语言的sync包中提供了一个开箱即用的并发安全版map-sync.Map
		不用像内置的map一样使用make函数初始化就能直接使用
		内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法

	原子操作
		代码中的加锁操作因为涉及到内核态的上下文切换会比较耗时、代价比较高
		针对基本数据类型我们还可以使用原子操作来保证并发安全，性能比加锁操作好
*/
// var m = sync.Map{}
var x int64
var l sync.Mutex
var wg sync.WaitGroup

func main() {
	// wg := sync.WaitGroup{}
	// for i := 0; i < 20; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		key := strconv.Itoa(n)
	// 		m.Store(key, n)
	// 		fmt.Printf("k=%v,v=%v\n", key, n)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	//一个示例比较一下互斥锁和原子操作的性能
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		// go add()// 普通版add函数 不是并发安全的
		// go mutexAdd()// 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
		//比较之后，后两者时间上感觉没啥大的区别
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

//普通版加函数
func add() {
	x++
	wg.Done()
}

//互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

//原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}
