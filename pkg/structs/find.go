package structs

import (
	"reflect"
)

func FindFieldByType[T any](v reflect.Value) (reflect.Value, bool) {
	if v.Kind() == reflect.Pointer && v.IsNil() {
		return reflect.Value{}, false
	}
	v = reflect.Indirect(v)
	if v.Type() == reflect.TypeFor[T]() {
		return v, true
	}
	switch v.Kind() {
	case reflect.Struct:
		for _, field := range v.Fields() {
			res, found := FindFieldByType[T](field)
			if found {
				return res, true
			}
		}
	case reflect.Slice:
		for i := range v.Len() {
			res, found := FindFieldByType[T](v.Index(i))
			if found {
				return res, true
			}
		}
	case reflect.Interface:
		return FindFieldByType[T](v.Elem())
	}
	return reflect.Value{}, false
}
