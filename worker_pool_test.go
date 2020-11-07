package gokit

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Score struct {
	Num int
}

func (s *Score) Do() {
	fmt.Println("num:", s.Num)
	time.Sleep(1 * time.Second)
}
func ExampleNewWorkerPool() {
	fmt.Println("=============")
	num := 3
	// debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	p := NewWorkerPool(ctx, num, wg)
	p.Run()
	// datanum := 100
	// go func() {
	// 	for i := 1; i <= datanum; i++ {
	// 		sc := &Score{Num: i}
	// 		p.AddJob(sc)
	// 	}
	// }()
	cancel()
	wg.Wait()
	// Output: Hello

}
