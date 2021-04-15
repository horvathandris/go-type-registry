package go_type_registry

import (
	"fmt"
	"reflect"
	"strings"
)

type TypeRegistry struct {
	m map[string]reflect.Type
}

func newTypeRegistry() TypeRegistry {
	return TypeRegistry{m: make(map[string]reflect.Type)}
}

func (registry TypeRegistry) set(s string, t reflect.Type) {
	registry.m[strings.ToLower(s)] = t
}

func (registry TypeRegistry) get(s string) (t reflect.Type, ok bool) {
	t, ok = registry.m[strings.ToLower(s)]
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
