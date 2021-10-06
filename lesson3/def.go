package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	//定义一个slice
	queue []int
	cond  *sync.Cond
}

// Push 往slice中插入数据
func (q *Queue) Push(i, num int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	//在queue这个slice中加入i
	q.queue = append(q.queue, i)
	fmt.Printf(">>>>>生产者%d创建了一个消息\n", num)

	q.cond.Broadcast() //通知所有等待的线程
	//q.cond.Signal()//通知某一个等待的线程
}

// Pop 从slice中取出数据
func (q *Queue) Pop(num int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	//如果slice中没有数据线程就等待
	for len(q.queue) == 0 {
		fmt.Println("I'm waiting 1...")
		q.cond.Wait()
		fmt.Println("I'm waiting 2...")
	}
	//取出slice中第一个数
	i := q.queue[0]
	//剩下的slice再赋值给slice
	q.queue = q.queue[1:]
	fmt.Printf("<<<<<消费者%d消费了%d\n", num, i)
}
