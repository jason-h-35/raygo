package tracer

import (
	"math"
)

func (m Mat[Size4]) Translate(x, y, z float64) Mat[Size4] {
	f := []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}

func (m Mat[Size4]) Scale(x, y, z float64) Mat[Size4] {
	f := []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}

func (m Mat[Size4]) RotateX(rad float64) Mat[Size4] {
	f := []float64{
		1, 0, 0, 0,
		0, math.Cos(rad), -math.Sin(rad), 0,
		0, math.Sin(rad), math.Cos(rad), 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}

func (m Mat[Size4]) RotateY(rad float64) Mat[Size4] {
	f := []float64{
		math.Cos(rad), 0, math.Sin(rad), 0,
		0, 1, 0, 0,
		-math.Sin(rad), 0, math.Cos(rad), 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}

func (m Mat[Size4]) RotateZ(rad float64) Mat[Size4] {
	f := []float64{
		math.Cos(rad), -math.Sin(rad), 0, 0,
		math.Sin(rad), math.Cos(rad), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}

func (m Mat[Size4]) Shear(xy, xz, yx, yz, zx, zy float64) Mat[Size4] {
	f := []float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}
