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
		{"", I4.Translate(5, -3, 2), NewPoint(-3, 4, 5), NewPoint(2, 1, 7)},
		{"", I4.Translate(5, -3, 2).Inverse(), NewPoint(-3, 4, 5), NewPoint(-8, 7, 3)},
		{"", I4.Translate(5, -3, 2), NewVector(-3, 4, 5), NewVector(-3, 4, 5)},
		{"", I4.Scale(2, 3, 4), NewPoint(-4, 6, 8), NewPoint(-8, 18, 32)},
		{"", I4.Scale(2, 3, 4), NewVector(-4, 6, 8), NewVector(-8, 18, 32)},
		{"", I4.Scale(2, 3, 4).Inverse(), NewVector(-4, 6, 8), NewVector(-2, 2, 2)},
		{"", I4.Scale(-1, 1, 1), NewPoint(2, 3, 4), NewPoint(-2, 3, 4)},
		{"", I4.RotateX(math.Pi / 4.), NewPoint(0, 1, 0), NewPoint(0, math.Sqrt(2)/2., math.Sqrt(2)/2.)},
		{"", I4.RotateX(math.Pi / 2.), NewPoint(0, 1, 0), NewPoint(0, 0, 1)},
		{"", I4.RotateX(math.Pi / 4.).Inverse(), NewPoint(0, 1, 0), NewPoint(0, math.Sqrt(2)/2., -math.Sqrt(2)/2.)},
		{"", I4.RotateY(math.Pi / 4.), NewPoint(0, 0, 1), NewPoint(math.Sqrt(2)/2., 0, math.Sqrt(2)/2.)},
		{"", I4.RotateY(math.Pi / 2.), NewPoint(0, 0, 1), NewPoint(1, 0, 0)},
		{"", I4.RotateZ(math.Pi / 4.), NewPoint(0, 1, 0), NewPoint(-math.Sqrt(2)/2., math.Sqrt(2)/2., 0)},
		{"", I4.RotateZ(math.Pi / 2.), NewPoint(0, 1, 0), NewPoint(-1, 0, 0)},
		{"", I4.Shear(1, 0, 0, 0, 0, 0), NewPoint(2, 3, 4), NewPoint(5, 3, 4)},
		{"", I4.Shear(0, 1, 0, 0, 0, 0), NewPoint(2, 3, 4), NewPoint(6, 3, 4)},
		{"", I4.Shear(0, 0, 1, 0, 0, 0), NewPoint(2, 3, 4), NewPoint(2, 5, 4)},
		{"", I4.Shear(0, 0, 0, 1, 0, 0), NewPoint(2, 3, 4), NewPoint(2, 7, 4)},
		{"", I4.Shear(0, 0, 0, 0, 1, 0), NewPoint(2, 3, 4), NewPoint(2, 3, 6)},
		{"", I4.Shear(0, 0, 0, 0, 0, 1), NewPoint(2, 3, 4), NewPoint(2, 3, 7)},
		{"", I4.RotateX(math.Pi / 2), NewPoint(1, 0, 1), NewPoint(1, -1, 0)},
		{"", I4.Scale(5, 5, 5), NewPoint(1, -1, 0), NewPoint(5, -5, 0)},
		{"", I4.Translate(10, 5, 7), NewPoint(5, -5, 0), NewPoint(15, 0, 7)},
		{"", I4.RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7), NewPoint(1, 0, 1), NewPoint(15, 0, 7)},
		{"", I4.Translate(10, 5, 7).Times(I4.Scale(5, 5, 5)).Times(I4.RotateX(math.Pi / 2)), NewPoint(1, 0, 1), NewPoint(15, 0, 7)},
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
