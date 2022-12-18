package attr

import "testing"

func BenchmarkAttr(b *testing.B) {
	var attr Attr

	for i := 0; i < b.N; i++ {
		attribute := New("attr", "attr")
		attr = attribute
	}

	_ = attr
}

func BenchmarkMapSimple(b *testing.B) {
	var attrMap map[string]any
	var input = Attrs{
		New("string", "string"),
		New("int", 0),
		New("float", 1.0),
	}

	for i := 0; i < b.N; i++ {
		attributeMap := Map(input...)
		attrMap = attributeMap
	}
	_ = attrMap
}

func BenchmarkMapComplex(b *testing.B) {
	var attrMap map[string]any
	var input = Attrs{
		New("string", "string"),
		New("int", 0),
		New("float", 1.0),
		New("complex", -0.3+0.6i),
		New("bool", true),
		New("string", "more"),
		New("custom_type", struct{ name string }{name: "nope"}),
	}

	for i := 0; i < b.N; i++ {
		attributeMap := Map(input...)
		attrMap = attributeMap
	}
	_ = attrMap
}
