package tracer

import (
	"math"
	"testing"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name      string
		transform Mat[Size4]
		pt        Tuple
		expect    Tuple
	}{
		{"Translation moves a point", I4.Translate(5, -3, 2), NewPoint(-3, 4, 5), NewPoint(2, 1, 7)},
		{"Inverse translation moves a point in the opposite direction", I4.Translate(5, -3, 2).Inverse(), NewPoint(-3, 4, 5), NewPoint(-8, 7, 3)},
		{"Translation does not affect vectors", I4.Translate(5, -3, 2), NewVector(-3, 4, 5), NewVector(-3, 4, 5)},
		{"Scaling matrix applied to a point", I4.Scale(2, 3, 4), NewPoint(-4, 6, 8), NewPoint(-8, 18, 32)},
		{"Scaling matrix applied to a vector", I4.Scale(2, 3, 4), NewVector(-4, 6, 8), NewVector(-8, 18, 32)},
		{"Inverse of scaling matrix", I4.Scale(2, 3, 4).Inverse(), NewVector(-4, 6, 8), NewVector(-2, 2, 2)},
		{"Reflection is scaling by a negative value", I4.Scale(-1, 1, 1), NewPoint(2, 3, 4), NewPoint(-2, 3, 4)},
		{"Rotating a point around X axis by quarter turn", I4.RotateX(math.Pi / 4.), NewPoint(0, 1, 0), NewPoint(0, math.Sqrt(2)/2., math.Sqrt(2)/2.)},
		{"Rotating a point around X axis by half turn", I4.RotateX(math.Pi / 2.), NewPoint(0, 1, 0), NewPoint(0, 0, 1)},
		{"Inverse of X rotation rotates in opposite direction", I4.RotateX(math.Pi / 4.).Inverse(), NewPoint(0, 1, 0), NewPoint(0, math.Sqrt(2)/2., -math.Sqrt(2)/2.)},
		{"Rotating a point around Y axis by quarter turn", I4.RotateY(math.Pi / 4.), NewPoint(0, 0, 1), NewPoint(math.Sqrt(2)/2., 0, math.Sqrt(2)/2.)},
		{"Rotating a point around Y axis by half turn", I4.RotateY(math.Pi / 2.), NewPoint(0, 0, 1), NewPoint(1, 0, 0)},
		{"Rotating a point around Z axis by quarter turn", I4.RotateZ(math.Pi / 4.), NewPoint(0, 1, 0), NewPoint(-math.Sqrt(2)/2., math.Sqrt(2)/2., 0)},
		{"Rotating a point around Z axis by half turn", I4.RotateZ(math.Pi / 2.), NewPoint(0, 1, 0), NewPoint(-1, 0, 0)},
		{"Shearing transformation moves x in proportion to y", I4.Shear(1, 0, 0, 0, 0, 0), NewPoint(2, 3, 4), NewPoint(5, 3, 4)},
		{"Shearing transformation moves x in proportion to z", I4.Shear(0, 1, 0, 0, 0, 0), NewPoint(2, 3, 4), NewPoint(6, 3, 4)},
		{"Shearing transformation moves y in proportion to x", I4.Shear(0, 0, 1, 0, 0, 0), NewPoint(2, 3, 4), NewPoint(2, 5, 4)},
		{"Shearing transformation moves y in proportion to z", I4.Shear(0, 0, 0, 1, 0, 0), NewPoint(2, 3, 4), NewPoint(2, 7, 4)},
		{"Shearing transformation moves z in proportion to x", I4.Shear(0, 0, 0, 0, 1, 0), NewPoint(2, 3, 4), NewPoint(2, 3, 6)},
		{"Shearing transformation moves z in proportion to y", I4.Shear(0, 0, 0, 0, 0, 1), NewPoint(2, 3, 4), NewPoint(2, 3, 7)},
		{"Individual transformations in sequence: rotation first", I4.RotateX(math.Pi / 2), NewPoint(1, 0, 1), NewPoint(1, -1, 0)},
		{"Individual transformations in sequence: scaling second", I4.Scale(5, 5, 5), NewPoint(1, -1, 0), NewPoint(5, -5, 0)},
		{"Individual transformations in sequence: translation last", I4.Translate(10, 5, 7), NewPoint(5, -5, 0), NewPoint(15, 0, 7)},
		{"Chained transformations must be applied in reverse order", I4.RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7), NewPoint(1, 0, 1), NewPoint(15, 0, 7)},
		{"Chained transformations in expanded form", I4.Translate(10, 5, 7).Times(I4.Scale(5, 5, 5)).Times(I4.RotateX(math.Pi / 2)), NewPoint(1, 0, 1), NewPoint(15, 0, 7)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.transform.TimesTuple(tt.pt)
			if !result.Equals(tt.expect) {
				t.Errorf("%v should have transformed to %v but was %v", tt.pt, tt.expect, result)
			}
		})
	}
}
