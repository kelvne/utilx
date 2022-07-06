package reflx

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// StructSetValue sets a value to a field on a struct
func StructSetValue(s interface{}, key string, v interface{}) error {
	value := IndirectValue(s)

	if value.Kind() != reflect.Struct {
		return errors.New("given value not a struct")
	}

	field := value.FieldByName(key)

	if reflect.ValueOf(field).IsZero() {
		return fmt.Errorf(
			"%s field not found for %s",
			field,
			reflect.TypeOf(value.Interface()).Name(),
		)
	}

	field.Set(reflect.ValueOf(v))

	return nil
}

// StructGetValue retrieves a field's value from a struct
func StructGetValue(s interface{}, key string) (interface{}, error) {
	value := IndirectValue(s)

	if value.Kind() != reflect.Struct {
		return nil, errors.New("given value not a struct")
	}

	field := value.FieldByName(key)

	if reflect.ValueOf(field).IsZero() {
		return nil, fmt.Errorf("field %s not found", key)
	}

	return field.Interface(), nil
}

// StructFieldsValuesNonZeroMap retrieves a map of all non zeroed fields from a struct
func StructFieldsValuesNonZeroMap(s interface{}) (map[string]interface{}, error) {
	value := IndirectValue(s)

	if value.Kind() != reflect.Struct {
		return nil, errors.New("given value not a struct")
	}

	m := make(map[string]interface{})

	for i := 0; i < value.NumField(); i += 1 {
		field := value.Field(i)
		if !field.IsZero() {
			fieldType := value.Type().Field(i)
			m[fieldType.Name] = field.Interface()
		}
	}

	return m, nil
}

// StructTagMap returns a map of field tag / name
func StructTagMap(v interface{}, tag string) (map[string]string, error) {
	value := IndirectValue(v)

	if value.Kind() != reflect.Struct {
		return nil, errors.New("given value not a struct")
	}

	m := make(map[string]string)

	for i := 0; i < value.NumField(); i += 1 {
		field := value.Type().Field(i)
		t := strings.Split(field.Tag.Get(tag), ",")[0]
		if strings.TrimSpace(t) != "" {
			m[t] = field.Name
		}
	}

	return m, nil
}
