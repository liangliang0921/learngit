package main

import (
	"fmt"
	"os"
)

// CreateFile 创建文件
func CreateFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// WriteFile 写文件
func WriteFile(f *os.File) {
	fmt.Println("wrirting")
	fmt.Fprintln(f, "data")
}

// CloseFile 关闭文件
func CloseFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	f := CreateFile("/tmp/defer.txt")
	defer CloseFile(f)
	WriteFile(f)
}
