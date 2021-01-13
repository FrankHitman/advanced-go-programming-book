package main

import (
	"fmt"
)

// func main() {
// 	done := make(chan int)
// 	go func() {
// 		fmt.Println("你好, 世界")
// 		aa := <-done
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(aa)
// 			time.Sleep(1*time.Second)
// 		}
// 	}()
// 	done <- 1
// }

// 你好, 世界
// 1
// 没有循环与没有sleep时候，可以打印出1，主线程结束，自线程是不是应该也结束。
// 有循环与没有sleep时候，可以打印10次1，在主进程结束的瞬间，子进程执行掉循环
// 有循环与有sleep时候，可以打印1次1，因为主线程结束，子线程也跟着结束

// func main()  {
// 	done := make(chan int)
//
// 	go func() {
// 		fmt.Println("hello, world")
// 		<- done
// 	}()
//
// 	done <- 1
// }

// “根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前”
// 所以以上可以打印出 hello, world

// func main()  {
// 	done := make(chan int, 1)
//
// 	go func() {
// 		fmt.Println("hello, world")
// 		<- done
// 	}()
//
// 	done <- 1
// }

// 以上打印不出来hello, world，由于是有缓冲的channel，那么主线程不会被block住，执行完就结束，
// 子线程没有机会执行

func main()  {
	done := make(chan int, 1)

	go func() {
		fmt.Println("hello, world")
		done <- 1
	}()
	<- done
}

// 对调对有缓冲的channel的读写顺序，以上又可以打印 hello, world 了
// “对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前”
