package tracer

import (
	"testing"
)

func TestTranslation(t *testing.T) {
	transform := I4.Translate(5, -3, 2)
	p := Point(-3, 4, 5)
	result := transform.Times(p)
	expect := Point(2, 1, 7)
	if !result.Equal(expect) {
		t.Errorf("blah blah blah")
	}
}
