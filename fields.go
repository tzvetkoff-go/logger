package logger

import (
	"sort"
)

// Fields is just an alias for map[string]interface{} with a few convenience methods added.
type Fields map[string]interface{}

// Copy returns a copy of the map.
func (f Fields) Copy() Fields {
	result := make(Fields, len(f))
	result.Update(f)
	return result
}

// Update updates the map with keys & values from the other map.
func (f Fields) Update(updates Fields) {
	for k, v := range updates {
		f[k] = v
	}
}

// Keys returns the map keys sorted alphabetically.
func (f Fields) Keys() []string {
	result := []string{}
	for k := range f {
		result = append(result, k)
	}
	sort.Strings(result)

	return result
}
