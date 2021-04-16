/*
	Package registry
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package registry

var Template = `
package {{.Package}}

var Types = []interface{}{ {{ range $index, $element := .TypeNames}}
	{{$element}}{},{{end}}
}
`
