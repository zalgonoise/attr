package attr

import json "github.com/goccy/go-json"

// Ptr is a generic function to create an Attr from a pointer value
//
// Using a generic approach allows the Attr.WithValue method to be
// scoped with certain constraints for specific applications
func Ptr[T any](key string, value *T) Attr {
	if key == "" {
		return nil
	}
	return &ptrAttr[T]{
		key: key,
		ptr: value,
	}
}

type ptrAttr[T any] struct {
	key string
	ptr *T
}

// Key returns the string key of the attribute Attr
func (p *ptrAttr[T]) Key() string {
	return p.key
}

// Value returns the (any) value of the attribute Attr
func (p *ptrAttr[T]) Value() any {
	if p.ptr == nil {
		return nil
	}
	return *p.ptr
}

// WithKey returns a copy of this Attr, with key `key`
func (p *ptrAttr[T]) WithKey(key string) Attr {
	return Ptr(key, p.ptr)
}

// WithValue returns a copy of this Attr, with value `value`
//
// It must be the same type of the original Attr, otherwise returns
// nil
func (p *ptrAttr[T]) WithValue(value any) Attr {
	if value == nil {
		return nil
	}

	v, ok := (value).(*T)
	if !ok {
		return nil
	}
	return Ptr(p.key, v)
}

// MarshalJSON encodes the attribute as a JSON object (key-value pair)
func (p *ptrAttr[T]) MarshalJSON() ([]byte, error) {
	var kv = map[string]any{}
	switch v := p.Value().(type) {
	case []Attr:
		kv[p.Key()] = mapAttrs(v...)
	case Attr:
		kv[p.Key()] = mapAttrs(v)
	default:
		kv[p.Key()] = p.Value()
	}

	return json.Marshal(kv)
}

func (p *ptrAttr[T]) MarshalText() (text []byte, err error) {
	return p.MarshalJSON()
}

// String implements fmt.Stringer
func (p *ptrAttr[T]) String() string {
	b, _ := p.MarshalJSON()
	return string(b)
}
