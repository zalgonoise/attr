package attr

import json "github.com/goccy/go-json"

type Attrs []Attr

// MarshalJSON encodes the attributes as a JSON object (key-value pairs)
func (attrs Attrs) MarshalJSON() ([]byte, error) {
	var kv = map[string]any{}
	for _, a := range attrs {
		switch v := a.Value().(type) {
		case []Attr:
			kv[a.Key()] = mapAttrs(v...)
		case Attr:
			kv[a.Key()] = mapAttrs(v)
		default:
			kv[a.Key()] = a.Value()
		}
	}

	return json.Marshal(kv)
}

func (attrs Attrs) MarshalText() (text []byte, err error) {
	return attrs.MarshalJSON()
}

// String implements fmt.Stringer
func (attrs Attrs) String() string {
	b, _ := attrs.MarshalJSON()
	return string(b)
}
