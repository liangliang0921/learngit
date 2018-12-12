package main

import (
	"crypto/sha1"
	"fmt"
)

// cha1 生成二进制文件和或者文本块的短标识
func main() {
	s := "sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
