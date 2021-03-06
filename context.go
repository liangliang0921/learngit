//package main
//
//import (
//	"context"
//	"fmt"
//)
//
//func func1(ctx context.Context) {
//	ctx = context.WithValue(ctx, "k1", "v2")
//	func2(ctx)
//}
//
//func func2(ctx context.Context) {
//	fmt.Println(ctx.Value("k1").(string))
//}
//
//func main() {
//	ctx := context.Background()
//	func1(ctx)
//}

// context 超时处理
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func func1(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	respC := make(chan int)
	// 处理逻辑
	go func() {
		time.Sleep(time.Second * 5)
		respC <- 10
	}()
	// 取消机制
	select {
	case <-ctx.Done():
		fmt.Println("cancel")
		return errors.New("cancel")
	case r := <-respC:
		fmt.Println(r)
		return nil
	}
}

func main() {
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func1(ctx, wg)
	time.Sleep(time.Second * 2)
	// 触发取消
	cancel()
	// 等待goroutine退出
	wg.Wait()
}
