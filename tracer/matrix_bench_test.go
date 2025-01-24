package tracer

import (
	"testing"
)

func BenchmarkMatrixCreation(b *testing.B) {
	vals2 := []float64{1, 2, 3, 4}
	vals3 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	vals4 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.Run("NewMat2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat2 = NewMat[Size2](vals2)
		}
	})

	b.Run("NewMat3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat3 = NewMat[Size3](vals3)
		}
	})

	b.Run("NewMat4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat4 = NewMat[Size4](vals4)
		}
	})
}

func BenchmarkMatrixOperations(b *testing.B) {
	m2 := NewMat[Size2]([]float64{1, 2, 3, 4})
	m3 := NewMat[Size3]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m4 := NewMat[Size4]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	t := NewPoint(1, 2, 3)

	b.Run("Times/2x2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat2 = m2.Times(m2)
		}
	})

	b.Run("Times/3x3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat3 = m3.Times(m3)
		}
	})

	b.Run("Times/4x4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat4 = m4.Times(m4)
		}
	})

	b.Run("TimesTuple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.tuple.t = m4.TimesTuple(t)
		}
	})

	b.Run("Transpose/2x2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat2 = m2.Transpose()
		}
	})

	b.Run("Transpose/4x4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat4 = m4.Transpose()
		}
	})
}

func BenchmarkMatrixDeterminants(b *testing.B) {
	m2 := NewMat[Size2]([]float64{1, 2, 3, 4})
	m3 := NewMat[Size3]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m4 := NewMat[Size4]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	b.Run("Determinant/2x2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.float = m2.Determinant()
		}
	})

	b.Run("Determinant/3x3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.float = m3.Determinant()
		}
	})

	b.Run("Determinant/4x4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.float = m4.Determinant()
		}
	})
}

func BenchmarkMatrixInverse(b *testing.B) {
	// Use matrices that we know are invertible
	m2 := NewMat[Size2]([]float64{4, 7, 2, 6})
	m3 := NewMat[Size3]([]float64{1, 2, 3, 0, 1, 4, 5, 6, 0})
	m4 := NewMat[Size4]([]float64{
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4,
	})

	b.Run("CanInverse/2x2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.bool = m2.CanInverse()
		}
	})

	// BUG: Minor only supported for Mat3 and Mat4
	// b.Run("Inverse/2x2", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		results.matrix.mat2 = m2.Inverse()
	// 	}
	// })

	b.Run("Inverse/3x3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat3 = m3.Inverse()
		}
	})

	b.Run("Inverse/4x4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat4 = m4.Inverse()
		}
	})
}

func BenchmarkMatrixSubOperations(b *testing.B) {
	m3 := NewMat[Size3]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m4 := NewMat[Size4]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	b.Run("SubMat/3x3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat2 = SubMat[Size3, Size2](m3, 0, 0)
		}
	})

	b.Run("SubMat/4x4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.matrix.mat3 = SubMat[Size4, Size3](m4, 0, 0)
		}
	})

	b.Run("Minor/3x3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.float = Minor(m3, 0, 0)
		}
	})

	b.Run("Cofactor/3x3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.float = Cofactor(m3, 0, 0)
		}
	})
}
