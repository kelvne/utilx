package utilx

import "strconv"

// Map is a more complete map
type Map map[string]interface{}

// String returns property value as string
func (m Map) String(key string, or ...string) string {
	if prop, ok := m[key]; ok {
		if value, ok := prop.(string); ok {
			return GetStringOrDefault("", value)
		}
	}
	return GetStringOrDefault("", or...)
}

// Int returns property value as int
func (m Map) Int(key string, or ...int) int {
	if prop, ok := m[key]; ok {
		switch value := prop.(type) {
		case string:
			if v, err := strconv.Atoi(value); err == nil {
				return GetNumberOrDefault(0, v)
			}
		case int:
			return value
		case uint:
		case uint32:
		case uint64:
		case int64:
		case int32:
		case float64:
		case float32:
			return int(value)
		}
	}
	return GetNumberOrDefault(0, or...)
}

// Bool returns property value as bool
func (m Map) Bool(key string, or ...bool) bool {
	if prop, ok := m[key]; ok {
		if value, ok := prop.(bool); ok {
			return GetBoolOrDefault(false, value)
		}
	}
	return GetBoolOrDefault(false, or...)
}
