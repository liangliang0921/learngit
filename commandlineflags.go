package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("wrod", "foo", "a string")

	numPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string

	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("wrod:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
