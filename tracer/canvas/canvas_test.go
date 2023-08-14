package canvas

import (
	"testing"
)

func Test_NewColor(t *testing.T) {
	c1 := Color{0, 0.5, 1}
	c2 := NewColor(0, 0.5, 1)
}
