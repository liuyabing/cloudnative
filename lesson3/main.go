package main

import (
	"sync"
	"time"
)

func main() {
	//初始化一个Queue这个struct的-指针对象,
	q := &Queue{
		queue: []int{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	//定义5个线程,每个线程通过q对象的Push方法放一个数
	for num := 1; num < 5; num++ {
		go func(n int) {
			q.Push(8, n)
			time.Sleep(2 * time.Second)
		}(num)
	}

	//定义5个线程,每个线程通过q对象的Pop方法取一个数
	for num := 1; num < 5; num++ {
		go func(n int) {
			for {
				q.Pop(n)
				time.Sleep(1 * time.Second)
			}
		}(num)
	}

	//永远阻塞, 让main函数不退出, 让它在后台一直执行
	select {}
}
