/*
	Package registry
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package registry

var Template = `
package {{.Package}}

import (
	"github.com/horvathandris/go-type-registry/registry"
	"reflect"
)

var TypeRegistry = registry.TypeRegistry{
	M: map[string]reflect.Type{ {{ range $index, $element := .TypeNames}}
		"{{$element | ToLower}}": reflect.TypeOf({{$element}}{}),{{end}}
	},
}
`
