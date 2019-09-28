package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup
var s1 = make([]int, 1)
var s2 = make([]int, 1)

func add1(i int) {
	time.Sleep(1 * time.Second)
	s1 = append(s1, i)
}

func add2(i int) {
	time.Sleep(1 * time.Second)
	s2 = append(s2, i)
	wg.Done()
}

func main() {
	s1[0], s2[0] = 0, 0
	fmt.Println("切片1:", s1, "切片2:", s2)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println("协程1:", i, "协程2:", i)
		go add1(i)
		go add2(i)
	}

	wg.Wait()

	b := []int{1, 3, 5, 2, 4, 6}
	fmt.Println("before sort:", b)
	sort.Ints(b)
	fmt.Println("after sort:", b)
	//time.Sleep(2 * time.Second)
	sort.Ints(s1)
	sort.Ints(s2)
	fmt.Println("切片1:", s1, "切片2:", s2)
}
