/*
	Package registry
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package registry

import (
	"fmt"
	"reflect"
	"strings"
)

type TypeRegistry struct {
	M map[string]reflect.Type
}

func newTypeRegistry() TypeRegistry {
	return TypeRegistry{M: make(map[string]reflect.Type)}
}

func (registry TypeRegistry) set(s string, t reflect.Type) {
	registry.M[strings.ToLower(s)] = t
}

func (registry TypeRegistry) get(s string) (t reflect.Type, ok bool) {
	t, ok = registry.M[strings.ToLower(s)]
	return
}

func Create(types []interface{}) TypeRegistry {
	var registry = newTypeRegistry()
	for _, t := range types {
		registry.set(strings.Split(fmt.Sprintf("%T", t), ".")[1], reflect.TypeOf(t))
	}
	return registry
}

func (registry TypeRegistry) MakeInstance(name string) interface{} {
	t, _ := registry.get(name)
	v := reflect.New(t).Elem()
	return v.Interface()
}
