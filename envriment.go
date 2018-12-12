package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))
	fmt.Println()

	// os.Environ() 列出所有环境变量的键值对
	for _, e := range os.Environ() {
		// strings.Split 打印所有的键
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
