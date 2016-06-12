package utils

// String lookups for an interface
func String(val interface{}) string {
	if value, ok := val.(string); ok {
		return value
	}
	return ""
}

// Bool lookups for an interface
func Bool(val interface{}) bool {
	if value, ok := val.(bool); ok {
		return value
	}
	return false
}

// Int lookups for an interface
func Int(val interface{}) int {
	if value, ok := val.(int); ok {
		return value
	}
	return 0
}

// Float64 lookups for an interface
func Float64(val interface{}) float64 {
	if value, ok := val.(float64); ok {
		return value
	}
	return 0.0
}

// Uint64 lookups for an interface
func Uint64(val interface{}) uint64 {
	if value, ok := val.(uint64); ok {
		return value
	}
	return 0
}
