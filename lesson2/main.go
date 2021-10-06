package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 10)
	timeout := make(chan bool)

	defer close(channel)

	//消费者
	go func() {
		//下面是创建了一个定时器
		ticker := time.NewTicker(time.Second)
		//每一秒定时器触发一次
		for range ticker.C {
			select {
			case <-timeout:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-channel)
			}
		}
	}()

	//生产者
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("channel <- ", i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 10)
	close(timeout)
	time.Sleep(time.Second * 1)
	fmt.Println("main process exit!")

}
