package easy_sql

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func assignPtrType(field reflect.Value) (any, error) {
	switch field.Type().String() {
	case "int":
		return (*int)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "bool":
		return (*bool)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "string":
		return (*string)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "time.Time":
		return (*time.Time)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "any":
		return (*any)(unsafe.Pointer(field.UnsafeAddr())), nil
	default:
		return nil, fmt.Errorf("type %s unreconized", field.Type().String())
	}
}

func GetPointers(v any) ([]any, error) {
	var vPointers []any
	var rv = reflect.ValueOf(v).Elem()
	for i := range rv.NumField() {
		ptr, err := assignPtrType(rv.Field(i))
		if err != nil {
			return nil, err
		}
		vPointers = append(vPointers, ptr)
	}
	return vPointers, nil
}
