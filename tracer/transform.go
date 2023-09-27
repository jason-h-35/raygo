package tracer

import (
	"math"
)

func (m Mat4) Translate(x, y, z float64) Mat4 {
	f := []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
	T := NewMat4(f)
	return T.Times(m)
}

func (m Mat4) Scale(x, y, z float64) Mat4 {
	f := []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
	T := NewMat4(f)
	return T.Times(m)
}

func (m Mat4) RotateX(rad float64) Mat4 {
	f := []float64{
		1, 0, 0, 0,
		0, math.Cos(rad), -math.Sin(rad), 0,
		0, math.Sin(rad), math.Cos(rad), 0,
		0, 0, 0, 1,
	}
	T := NewMat4(f)
	return T.Times(m)
}

func (m Mat4) RotateY(rad float64) Mat4 {
	f := []float64{
		math.Cos(rad), 0, math.Sin(rad), 0,
		0, 1, 0, 0,
		-math.Sin(rad), 0, math.Cos(rad), 0,
		0, 0, 0, 1,
	}
	T := NewMat4(f)
	return T.Times(m)
}

func (m Mat4) RotateZ(rad float64) Mat4 {
	f := []float64{
		math.Cos(rad), -math.Sin(rad), 0, 0,
		math.Sin(rad), math.Cos(rad), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	T := NewMat4(f)
	return T.Times(m)
}
