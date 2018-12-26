package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}

	// 建立一个模板,
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}

	// 将 struct 与模板融合，合成结果放在 os.stdout
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
	// 打印模板的名字
	fmt.Println(tmpl.Name())

	// 如果一个文件的模本内容在一个文件（a.txt）中，将模板的内容发在一个文件里面
	// 建立模板
	// tmpl,err := template.ParseFiles("a.txt")
	// 当文件模板太多的时候，可以使用下面的写法，可以写多个文件
	// tmpl,err := template.ParseGlob("*.txt")

	// 当模板太多的时候可以使指定模板，比如下面指定sweaters模板
	// err := tmpl.ExecuteTemplate(os.Stdout,"test",sweaters)
}
