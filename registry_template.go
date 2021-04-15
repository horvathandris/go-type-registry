package go_type_registry

var registryTemplate = `
package go_type_registry

var Types = []interface{}{ {{ range $index, $element := .}}
	{{$element}}{},{{end}}
}
`
