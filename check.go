package easy_sql

import "reflect"

func IsPtr(rv reflect.Value) bool {
	return rv.Kind() == reflect.Pointer
}

func IsStruct(rv reflect.Value) bool {
	return rv.Elem().Kind() == reflect.Struct
}
