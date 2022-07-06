package reflx

import (
	"errors"
	"reflect"
)

// SliceOf returns a slice of the type of given value
func SliceOf(v interface{}, size int) interface{} {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(v)), 0, size)
	items := reflect.New(slice.Type())
	items.Elem().Set(slice)
	return items.Elem().Interface()
}

// SliceElem returns an empty interface of an element of a slice
func SliceElem(v interface{}) (interface{}, error) {
	slice := reflect.TypeOf(IndirectValue(v).Interface())
	if slice.Kind() != reflect.Slice {
		return nil, errors.New("value is not a slice")
	}
	return reflect.New(slice.Elem().Elem()).Interface(), nil
}
