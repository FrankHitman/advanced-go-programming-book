package main

import (
	"fmt"
	"sync"
	"time"
)

// 利用select的随机选择管道性，实现生成随机数序列
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for {
// 			select {
// 			case ch <- 0:
// 			case ch <- 1:
// 			}
// 		}
// 	}()
// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }

// // 实现goroutine的退出控制，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前
// func worker(cancel chan bool) {
// 	for {
// 		select {
// 		default:
// 			fmt.Println("hello")
// 			// 正常工作
// 		case <-cancel:
// 			// 退出
// 		}
// 	}
// }
//
// func main() {
// 	cancel := make(chan bool)
// 	go worker(cancel)
//
// 	time.Sleep(time.Second)
// 	cancel <- true
// }
// // 如上如果注释掉 case <-cancel:就会导致程序停不下来

// // 关闭管道来实现广播，所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标志
// func worker(cannel chan bool) {
// 	for {
// 		select {
// 		default:
// 			fmt.Println("hello")
// 			// 正常工作
// 		case <-cannel:
// 			fmt.Println("goodbye")
// 			// 退出
// 		}
// 	}
// }
//
// func main() {
// 	cancel := make(chan bool)
// 	for i := 0; i < 10; i++ {
// 		go worker(cancel)
// 	}
// 	time.Sleep(time.Second)
// 	close(cancel)
// }
// // 如上执行多次，打印goodbye次数不一定是10次，有11次，13次，0次。所以主线程关闭channel操作，不能保证子线程把清理工作完成

// 使用waitgroup同步主线程与子线程的操作
func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cannel:
			fmt.Println("goodbye")
			return
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
// 以上可以保证每次执行都是打印10次goodbye

