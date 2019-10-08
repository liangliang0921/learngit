package main

import (
	"fmt"
	"time"
)

func arrayAppend(array *[]int, i int) {
	*array = append(*array, i)
	fmt.Println(array)
}

func main() {
	array := []int{}
	array1 := []int{}
	for i := 0; i < 10; i++ {
		go arrayAppend(&array, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(array)

	for i := 0; i < 10; i++ {
		array1 = append(array1, i)
		fmt.Println("切片长度", cap(array1), "切片容量", len(array1))
	}
}
