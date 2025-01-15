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
		{
			name:      "Translation moves a point",
			transform: I4.Translate(5, -3, 2),
			pt:        NewPoint(-3, 4, 5),
			expect:    NewPoint(2, 1, 7),
		},
		{
			name:      "Inverse translation moves a point in the opposite direction",
			transform: I4.Translate(5, -3, 2).Inverse(),
			pt:        NewPoint(-3, 4, 5),
			expect:    NewPoint(-8, 7, 3),
		},
		{
			name:      "Translation does not affect vectors",
			transform: I4.Translate(5, -3, 2),
			pt:        NewVector(-3, 4, 5),
			expect:    NewVector(-3, 4, 5),
		},
		{
			name:      "Scaling matrix applied to a point",
			transform: I4.Scale(2, 3, 4),
			pt:        NewPoint(-4, 6, 8),
			expect:    NewPoint(-8, 18, 32),
		},
		{
			name:      "Scaling matrix applied to a vector",
			transform: I4.Scale(2, 3, 4),
			pt:        NewVector(-4, 6, 8),
			expect:    NewVector(-8, 18, 32),
		},
		{
			name:      "Inverse of scaling matrix",
			transform: I4.Scale(2, 3, 4).Inverse(),
			pt:        NewVector(-4, 6, 8),
			expect:    NewVector(-2, 2, 2),
		},
		{
			name:      "Reflection is scaling by a negative value",
			transform: I4.Scale(-1, 1, 1),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(-2, 3, 4),
		},
		{
			name:      "Rotating a point around X axis by quarter turn",
			transform: I4.RotateX(math.Pi / 4.),
			pt:        NewPoint(0, 1, 0),
			expect:    NewPoint(0, math.Sqrt(2)/2., math.Sqrt(2)/2.),
		},
		{
			name:      "Rotating a point around X axis by half turn",
			transform: I4.RotateX(math.Pi / 2.),
			pt:        NewPoint(0, 1, 0),
			expect:    NewPoint(0, 0, 1),
		},
		{
			name:      "Inverse of X rotation rotates in opposite direction",
			transform: I4.RotateX(math.Pi / 4.).Inverse(),
			pt:        NewPoint(0, 1, 0),
			expect:    NewPoint(0, math.Sqrt(2)/2., -math.Sqrt(2)/2.),
		},
		{
			name:      "Rotating a point around Y axis by quarter turn",
			transform: I4.RotateY(math.Pi / 4.),
			pt:        NewPoint(0, 0, 1),
			expect:    NewPoint(math.Sqrt(2)/2., 0, math.Sqrt(2)/2.),
		},
		{
			name:      "Rotating a point around Y axis by half turn",
			transform: I4.RotateY(math.Pi / 2.),
			pt:        NewPoint(0, 0, 1),
			expect:    NewPoint(1, 0, 0),
		},
		{
			name:      "Rotating a point around Z axis by quarter turn",
			transform: I4.RotateZ(math.Pi / 4.),
			pt:        NewPoint(0, 1, 0),
			expect:    NewPoint(-math.Sqrt(2)/2., math.Sqrt(2)/2., 0),
		},
		{
			name:      "Rotating a point around Z axis by half turn",
			transform: I4.RotateZ(math.Pi / 2.),
			pt:        NewPoint(0, 1, 0),
			expect:    NewPoint(-1, 0, 0),
		},
		{
			name:      "Shearing transformation moves x in proportion to y",
			transform: I4.Shear(1, 0, 0, 0, 0, 0),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(5, 3, 4),
		},
		{
			name:      "Shearing transformation moves x in proportion to z",
			transform: I4.Shear(0, 1, 0, 0, 0, 0),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(6, 3, 4),
		},
		{
			name:      "Shearing transformation moves y in proportion to x",
			transform: I4.Shear(0, 0, 1, 0, 0, 0),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(2, 5, 4),
		},
		{
			name:      "Shearing transformation moves y in proportion to z",
			transform: I4.Shear(0, 0, 0, 1, 0, 0),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(2, 7, 4),
		},
		{
			name:      "Shearing transformation moves z in proportion to x",
			transform: I4.Shear(0, 0, 0, 0, 1, 0),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(2, 3, 6),
		},
		{
			name:      "Shearing transformation moves z in proportion to y",
			transform: I4.Shear(0, 0, 0, 0, 0, 1),
			pt:        NewPoint(2, 3, 4),
			expect:    NewPoint(2, 3, 7),
		},
		{
			name:      "Individual transformations in sequence: rotation first",
			transform: I4.RotateX(math.Pi / 2),
			pt:        NewPoint(1, 0, 1),
			expect:    NewPoint(1, -1, 0),
		},
		{
			name:      "Individual transformations in sequence: scaling second",
			transform: I4.Scale(5, 5, 5),
			pt:        NewPoint(1, -1, 0),
			expect:    NewPoint(5, -5, 0),
		},
		{
			name:      "Individual transformations in sequence: translation last",
			transform: I4.Translate(10, 5, 7),
			pt:        NewPoint(5, -5, 0),
			expect:    NewPoint(15, 0, 7),
		},
		{
			name:      "Chained transformations must be applied in reverse order",
			transform: I4.RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7),
			pt:        NewPoint(1, 0, 1),
			expect:    NewPoint(15, 0, 7),
		},
		{
			name:      "Chained transformations in expanded form",
			transform: I4.Translate(10, 5, 7).Times(I4.Scale(5, 5, 5)).Times(I4.RotateX(math.Pi / 2)),
			pt:        NewPoint(1, 0, 1),
			expect:    NewPoint(15, 0, 7),
		},
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
