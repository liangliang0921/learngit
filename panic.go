package main

import (
	"fmt"
	"os"
	"time"
)

// panic 通常表示程序在正常运行的时候不应该出现的错误
func test1() {
	panic("a problem")

	_, err := os.Create("/tmp/panic.txt")
	if err != nil {
		panic(err)
	}
}

// 发生panic只会执行当前协程中的defer()
func test2() {
	defer fmt.Println("in main")
	go func() {
		defer fmt.Println("in goroutine")
		panic("")
	}()
}

func test3() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panic : %s\n", e)
		}
	}()
	index := 4
	arr := []int{1, 2, 3}
	_ = arr[index]
}

func main() {

	test3()

	time.Sleep(1 * time.Second)
}
