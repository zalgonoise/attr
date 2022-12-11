package attr

import json "github.com/goccy/go-json"

type Attrs []Attr

// MarshalJSON encodes the attributes as a JSON object (key-value pairs)
func (attrs Attrs) MarshalJSON() ([]byte, error) {
	return json.Marshal(Map(attrs...))
}

func (attrs Attrs) MarshalText() (text []byte, err error) {
	return attrs.MarshalJSON()
}

// String implements fmt.Stringer
func (attrs Attrs) String() string {
	b, _ := attrs.MarshalJSON()
	return string(b)
}
