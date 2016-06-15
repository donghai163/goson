package goson

import (
	"encoding/json"
	"strconv"
)

// Goson the helper go object
type Goson struct {
	i interface{}
}

/*
	bool, for JSON booleans
	float64, for JSON numbers
	string, for JSON strings
	[]interface{}, for JSON arrays
	map[string]interface{}, for JSON objects
	nil for JSON null

	switch t := g.i.(type) {
	case bool:
	case float64:
	case string:
	case []interface{}:
	case map[string]interface{}:
	default:
	}
*/

// Parse will create a goson object from json data
func Parse(data []byte) (*Goson, error) {
	var i interface{}
	if err := json.Unmarshal(data, &i); err != nil {
		return nil, err
	}
	return &Goson{i: i}, nil
}

// Get returns a goson object from a key
func (g *Goson) Get(key string) *Goson {
	if m, ok := g.i.(map[string]interface{}); ok {
		return &Goson{i: m[key]}
	}
	return &Goson{i: new(interface{})}
}

// Value will retrieve the underlying interface value.
func (g *Goson) Value() interface{} {
	return g.i
}

// Len will call len() on the underlying value
func (g *Goson) Len() int {
	switch t := g.i.(type) {
	case string:
		return len(t)
	case []interface{}:
		return len(t)
	case map[string]interface{}:
		return len(t)
	default:
		return 0
	}
}

// Index is a replacement for accessing the index of a goson object if an array.
func (g *Goson) Index(index int) *Goson {
	if v, ok := g.i.([]interface{}); ok {
		return &Goson{i: v[index]}
	}
	return nil
}

// Bool returns the bool value
func (g *Goson) Bool() bool {
	if v, ok := g.i.(bool); ok {
		return v
	}
	return false
}

// Int returns the underlying Int value.
func (g *Goson) Int() int {
	if v, ok := g.i.(float64); ok {
		return int(v)
	}
	return 0
}

// Float returns the underlying float64 value.
func (g *Goson) Float() float64 {
	if v, ok := g.i.(float64); ok {
		return v
	}
	return 0
}

// Slice returns the underlying slice value otherwise an empty slice.
func (g *Goson) Slice() []interface{} {
	if v, ok := g.i.([]interface{}); ok {
		return v
	}
	return []interface{}{}
}

// Map returns the underlying map otherwise an empty map.
func (g *Goson) Map() map[string]interface{} {
	if v, ok := g.i.(map[string]interface{}); ok {
		return v
	}
	return map[string]interface{}{}
}

// String will return the string version of any of the underlying objects.
func (g *Goson) String() string {
	switch t := g.i.(type) {
	case bool:
		return strconv.FormatBool(t)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		return t
	case []interface{}:
		data, err := json.MarshalIndent(t, "", "\t")
		if err == nil {
			return string(data)
		}
	case map[string]interface{}:
		data, err := json.MarshalIndent(t, "", "\t")
		if err == nil {
			return string(data)
		}
	}
	return ""
}
