package reflx

import "reflect"

// Clone clones a value
func Clone(v interface{}) interface{} {
	value := IndirectValue(v)
	cloned := reflect.New(value.Type())
	cloned.Elem().Set(value)
	return cloned.Elem().Interface()
}

// PtrTo returns the pointer to a value
func PtrTo(v interface{}) interface{} {
	value := IndirectValue(v)
	ptrTo := reflect.New(reflect.TypeOf(value.Interface()))
	ptrTo.Elem().Set(value)
	return ptrTo.Interface()
}

// IndirectValue returns the value of a variable or the value it points to
func IndirectValue(v interface{}) reflect.Value {
	if rv, ok := v.(reflect.Value); ok {
		for rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		return rv
	}
	return IndirectValue(reflect.ValueOf(v))
}
