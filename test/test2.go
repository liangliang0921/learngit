package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 产生随机字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)
	result := []byte{}

	// 因为 golang 如果每次的种子数是一样的，所以每次运行的结果都是一样的，所以一时间做为种子数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r)
	for i := 0; i < length; i++ {
		s := r.Intn(len(bytes))
		fmt.Println(s)
		result = append(result, bytes[s])
		fmt.Println(string(result))
	}
	return string(result)
}

//func Random(length int) string {
//	r := rand.New()
//}

func main() {
	fmt.Println(RandomStr(10))
}
