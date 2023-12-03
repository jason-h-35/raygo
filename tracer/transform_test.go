package tracer

import (
	"math"
	"testing"
)

func TestTransform(t *testing.T) {
	data := []struct {
		transform Mat4
		pt        Tuple
		expect    Tuple
	}{
		{I4.Translate(5, -3, 2), NewPointTuple(-3, 4, 5), NewPointTuple(2, 1, 7)},
		{I4.Translate(5, -3, 2).Inverse(), NewPointTuple(-3, 4, 5), NewPointTuple(-8, 7, 3)},
		{I4.Translate(5, -3, 2), NewVectorTuple(-3, 4, 5), NewVectorTuple(-3, 4, 5)},
		{I4.Scale(2, 3, 4), NewPointTuple(-4, 6, 8), NewPointTuple(-8, 18, 32)},
		{I4.Scale(2, 3, 4), NewVectorTuple(-4, 6, 8), NewVectorTuple(-8, 18, 32)},
		{I4.Scale(2, 3, 4).Inverse(), NewVectorTuple(-4, 6, 8), NewVectorTuple(-2, 2, 2)},
		{I4.Scale(-1, 1, 1), NewPointTuple(2, 3, 4), NewPointTuple(-2, 3, 4)},
		{I4.RotateX(math.Pi / 4.), NewPointTuple(0, 1, 0), NewPointTuple(0, math.Sqrt(2)/2., math.Sqrt(2)/2.)},
		{I4.RotateX(math.Pi / 2.), NewPointTuple(0, 1, 0), NewPointTuple(0, 0, 1)},
		{I4.RotateX(math.Pi / 4.).Inverse(), NewPointTuple(0, 1, 0), NewPointTuple(0, math.Sqrt(2)/2., -math.Sqrt(2)/2.)},
		{I4.RotateY(math.Pi / 4.), NewPointTuple(0, 0, 1), NewPointTuple(math.Sqrt(2)/2., 0, math.Sqrt(2)/2.)},
		{I4.RotateY(math.Pi / 2.), NewPointTuple(0, 0, 1), NewPointTuple(1, 0, 0)},
		{I4.RotateZ(math.Pi / 4.), NewPointTuple(0, 1, 0), NewPointTuple(-math.Sqrt(2)/2., math.Sqrt(2)/2., 0)},
		{I4.RotateZ(math.Pi / 2.), NewPointTuple(0, 1, 0), NewPointTuple(-1, 0, 0)},
		{I4.Shear(1, 0, 0, 0, 0, 0), NewPointTuple(2, 3, 4), NewPointTuple(5, 3, 4)},
		{I4.Shear(0, 1, 0, 0, 0, 0), NewPointTuple(2, 3, 4), NewPointTuple(6, 3, 4)},
		{I4.Shear(0, 0, 1, 0, 0, 0), NewPointTuple(2, 3, 4), NewPointTuple(2, 5, 4)},
		{I4.Shear(0, 0, 0, 1, 0, 0), NewPointTuple(2, 3, 4), NewPointTuple(2, 7, 4)},
		{I4.Shear(0, 0, 0, 0, 1, 0), NewPointTuple(2, 3, 4), NewPointTuple(2, 3, 6)},
		{I4.Shear(0, 0, 0, 0, 0, 1), NewPointTuple(2, 3, 4), NewPointTuple(2, 3, 7)},
		{I4.RotateX(math.Pi / 2), NewPointTuple(1, 0, 1), NewPointTuple(1, -1, 0)},
		{I4.Scale(5, 5, 5), NewPointTuple(1, -1, 0), NewPointTuple(5, -5, 0)},
		{I4.Translate(10, 5, 7), NewPointTuple(5, -5, 0), NewPointTuple(15, 0, 7)},
		{I4.RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7), NewPointTuple(1, 0, 1), NewPointTuple(15, 0, 7)},
		{I4.Translate(10, 5, 7).Times(I4.Scale(5, 5, 5)).Times(I4.RotateX(math.Pi / 2)), NewPointTuple(1, 0, 1), NewPointTuple(15, 0, 7)},
	}
	for _, row := range data {
		result := row.transform.TimesTuple(row.pt)
		if !result.Equals(row.expect) {
			t.Errorf("%v should have transformed to %v but was %v", row.pt, row.expect, result)
		}
	}
}
