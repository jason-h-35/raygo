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
	return NewMat[Size4](f).Times(m)
}

func (m Mat[Size4]) Scale(x, y, z float64) Mat[Size4] {
	f := []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}

func (m Mat[Size4]) RotateX(rad float64) Mat[Size4] {
	sin, cos := math.Sincos(rad)
	f := []float64{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f).Times(m)
	return T.Times(m)
}

func (m Mat[Size4]) RotateY(rad float64) Mat[Size4] {
	sin, cos := math.Sincos(rad)
	f := []float64{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1,
	}
	T := NewMat[Size4](f)
	return T.Times(m)
}

func (m Mat[Size4]) RotateZ(rad float64) Mat[Size4] {
	sin, cos := math.Sincos(rad)
	f := []float64{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
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
