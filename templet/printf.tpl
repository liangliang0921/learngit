// 模板

package {{ .Package  }}

import(
	"{{ .Import }}"
)

func {{ .Func }}(){
	{{ .Print }}("{{ .PrintText }}")
}

