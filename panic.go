package main

import (
	"os"
)

// panic 通常表示程序在正常运行的时候不应该出现的错误
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/panic.txt")
	if err != nil {
		panic(err)
	}
}
