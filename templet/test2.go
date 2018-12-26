package main

import (
	_ "fmt"
	"os"
	"text/template"
)

// 模板套用
func main() {
	muban1 := `hi,{{template "M2"}},hi,{{template "M2"}}`
	muban2 := `hi 我是模板2，{{template "M3"}}`
	muban3 := "haha 我是模板3"

	tmp, err := template.New("M1").Parse(muban1)
	if err != nil {
		panic(err)
	}

	tmp.New("M2").Parse(muban2)
	if err != nil {
		panic(err)
	}

	tmp.New("M3").Parse(muban3)
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
