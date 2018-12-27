package main

import (
	"os"
	"text/template"
)

type CodeStr struct {
	Package   string // 包名
	Import    string // 导入的包名
	Func      string // 函数名
	Print     string // 答应函数名
	PrintText string // 打印的文本
}

// 通过 filename 打印出一个go文件的名
func FileName(FileName string) string {
	s := FileName
	s += ".go"
	return s
}

// 通过本模板实现一个"打印任意字符串的 golang 代码"
func main() {
	// 根据文件创建一个模板,假设保存模板的名字就是 printf.tpl ，建立模板
	tmpl, err := template.ParseFiles("printf.tpl")
	if err != nil {
		panic(err)
	}

	sweaters := CodeStr{"main", "fmt", "main", "fmt.Println", "haha"}
	filename := FileName("Printhaha")

	// 打开并创建一个文件
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// 将struct 与模板融合
	err = tmpl.Execute(file, sweaters)
	if err != nil {
		panic(err)
	}
}
