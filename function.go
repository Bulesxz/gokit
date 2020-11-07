package gokit

import (
	"reflect"
)

/*
常用封装函数
*/

// InArray 元素在数组中(只支持简单类型)
func InArray(item interface{}, array interface{}) bool {
	itypeof := reflect.TypeOf(item)
	arrtypeof := reflect.TypeOf(array)
	if arrtypeof.Elem().Kind() != itypeof.Kind() {
		panic("type mismatch")
	}
	sVal := reflect.ValueOf(array)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == item {
				return true
			}
		}
		return false
	}
	panic("type mismatch")
}
