/*
	Package registry
	Copyright © 2021 András Horváth horvath.andras216@gmail.com
*/
package registry

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type TypeRegistry map[string]reflect.Type

func newTypeRegistry() TypeRegistry {
	return make(map[string]reflect.Type)
}

func (registry TypeRegistry) set(s string, t reflect.Type) {
	registry[strings.ToLower(s)] = t
}

func (registry TypeRegistry) get(s string) (t reflect.Type, ok bool) {
	t, ok = registry[strings.ToLower(s)]
	return
}

func Create(types []interface{}) TypeRegistry {
	var registry = newTypeRegistry()
	for _, t := range types {
		registry.set(strings.Split(fmt.Sprintf("%T", t), ".")[1], reflect.TypeOf(t))
	}
	return registry
}

func (registry TypeRegistry) MakeInstance(name string) (interface{}, error) {
	t, exists := registry.get(name)
	if !exists {
		return nil, errors.New("no such type in the registry")
	}
	v := reflect.New(t).Elem()
	return v.Interface(), nil
}
