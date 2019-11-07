package main

import (
	"fmt"
)

func main() {
	mm := make(map[string]interface{})

	if len(mm) == 0 {
		fmt.Println("map is nil")
	}
}
