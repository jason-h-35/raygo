package tracer

import (
	"fmt"
)

type Matrix[T int] interface {
	At(i, j int) float64
	Equals(other Matrix[T]) bool
	Determinant() float64
	CanInverse() bool
}

type Mat[T int] struct {
	vals [][]float64
	size T
}

type MatVal struct {
	i   int
	j   int
	val float64
}

type Size2 int
type Size3 int
type Size4 int

const (
	_2 Size2 = 2
	_3 Size3 = 3
	_4 Size4 = 4
)

var I2 = NewMat[Size2]([]float64{1, 0, 0, 1})
var I3 = NewMat[Size3]([]float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
var I4 = NewMat[Size4]([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

var Z2 = NewMat[Size2](make([]float64, 4))
var Z3 = NewMat[Size3](make([]float64, 9))
var Z4 = NewMat[Size4](make([]float64, 16))

func NewMatVal(i int, j int, val float64) MatVal {
	return MatVal{i, j, val}
}

func NewMat[T int](a []float64) Mat[T] {
	var size T
	expected := size * size
	if len(a) != int(expected) {
		panic(fmt.Sprintf("Mat%d needs %d elements", size, expected))
	}

	// Initialize the 2D slice
	vals := make([][]float64, int(size))
	for i := range vals {
		vals[i] = make([]float64, int(size))
	}

	// Fill the values
	for i := 0; i < int(size); i++ {
		for j := 0; j < int(size); j++ {
			vals[i][j] = a[i*int(size)+j]
		}
	}

	return Mat[T]{
		vals: vals,
		size: size,
	}
}

func (m Mat[T]) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m1 Mat[T]) Equals(m2 Mat[T]) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (a Mat[T]) Times(b Mat[T]) Mat[T] {
	var size T
	result := make([]float64, size*size)
	for i := 0; i < int(size); i++ {
		for j := 0; j < int(size); j++ {
			next := 0.0
			for k := 0; k < int(size); k++ {
				next += a.vals[i][k] * b.vals[k][j]
			}
			result[i*int(size)+j] = next
		}
	}
	return NewMat[T](result)
}

func (a Mat[T]) TimesTuple(b Tuple) Tuple {
	var size T
	if size != 4 {
		panic("TimesTuple only works with 4x4 matrices")
	}
	result := make([]float64, 4)
	bArr := b.AsArray()
	for i := 0; i < 4; i++ {
		next := 0.0
		for k := 0; k < 4; k++ {
			next += a.vals[i][k] * bArr[k]
		}
		result[i] = next
	}
	return NewTuple(result[0], result[1], result[2], result[3])
}

func (a Mat[T]) Transpose() Mat[T] {
	var size T
	result := make([]float64, size*size)
	for i := 0; i < int(size); i++ {
		for j := 0; j < int(size); j++ {
			result[j*int(size)+i] = a.vals[i][j]
		}
	}
	return NewMat[T](result)
}

func SubMat[T, U int](a Mat[T], is, js int) Mat[U] {
	s := make([]float64, 0)
	for i, row := range a.vals {
		for j, val := range row {
			if i != is && j != js {
				s = append(s, val)
			}
		}
	}
	return NewMat[U](s)
}

func Minor[T int](a Mat[T], is, js int) float64 {
	var size T
	switch size {
	case 3:
		return SubMat[int, int](a, is, js).Determinant()
	case 4:
		return SubMat[int, int](a, is, js).Determinant()
	default:
		panic("Minor only supported for Mat3 and Mat4")
	}
}

func Cofactor[T int](a Mat[T], is, js int) float64 {
	if (is+js)%2 == 1 {
		return -Minor(a, is, js)
	}
	return Minor(a, is, js)
}

func (a Mat[T]) Determinant() float64 {
	var size T
	switch size {
	case 2:
		return a.vals[0][0]*a.vals[1][1] - a.vals[0][1]*a.vals[1][0]
	case 3, 4:
		det, i := 0.0, 0
		for j, val := range a.vals[i] {
			det += val * Cofactor(a, i, j)
		}
		return det
	default:
		panic(fmt.Sprintf("Determinant not implemented for size %d", size))
	}

}

func (a Mat[T]) CanInverse() bool {
	return abs(a.Determinant()) > eps
}

func (a Mat[T]) Inverse() Mat[T] {
	if !a.CanInverse() {
		panic("can't inverse this matrix")
	}
	var size T
	inverse := NewMat[T](make([]float64, size*size))
	for i, row := range a.vals {
		for j := range row {
			inverse.vals[j][i] = Cofactor(a, i, j) / a.Determinant()
		}
	}
	return inverse
}
