package util

import (
	"fmt"
	"regexp"
	"strconv"
)

// MatcherLookup returns the map value of the named captures.
func MatcherLookup(match []string, re *regexp.Regexp) map[string]string {
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	return result
}

// MustAtoi converts the given string to an integer.
// It panics if the conversion fails.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("MustAtoi: %s", err))
	}
	return i
}

// MustBeString converts the given value to a string.
// It panics if the conversion fails.
func MustBeString(value interface{}) string {
	v, ok := value.(string)
	if !ok {
		panic(fmt.Errorf("MustString: underlying type of input interface must be a string."))
	}
	return v
}
