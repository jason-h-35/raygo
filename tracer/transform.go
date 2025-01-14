package tracer

import (
	"math"
)

// Transform functions apply geometric transformations to 4x4 matrices.
// Each transformation returns a new matrix without modifying the original.
// The transformations follow the right-hand rule and are applied in order
// from right to left when chained.

// Translate returns a new matrix translated by (x,y,z)
func (m Mat[Size4]) Translate(x, y, z float64) Mat[Size4] {
	f := []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}

// Scale returns a new matrix scaled by (x,y,z) factors
func (m Mat[Size4]) Scale(x, y, z float64) Mat[Size4] {
	f := []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}

// RotateX returns a new matrix rotated around the X axis by rad radians
func (m Mat[Size4]) RotateX(rad float64) Mat[Size4] {
	sin, cos := math.Sincos(rad)
	f := []float64{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}

// RotateY returns a new matrix rotated around the Y axis by rad radians
func (m Mat[Size4]) RotateY(rad float64) Mat[Size4] {
	sin, cos := math.Sincos(rad)
	f := []float64{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}

// RotateZ returns a new matrix rotated around the Z axis by rad radians
func (m Mat[Size4]) RotateZ(rad float64) Mat[Size4] {
	sin, cos := math.Sincos(rad)
	f := []float64{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}

// Shear returns a new matrix sheared by the given factors
// xy: X along Y, xz: X along Z
// yx: Y along X, yz: Y along Z
// zx: Z along X, zy: Z along Y
func (m Mat[Size4]) Shear(xy, xz, yx, yz, zx, zy float64) Mat[Size4] {
	f := []float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	}
	return NewMat[Size4](f).Times(m)
}
