package main

import (
	"fmt"
)

func main() {
	a := 1 >> 100
	fmt.Println(a)
	b := 100 << 1
	fmt.Println(b)
}
