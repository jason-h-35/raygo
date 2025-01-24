package tracer

import (
	"math"
	"testing"
)

// Prevent compiler optimizations
var resultMat Mat[Size4]

func BenchmarkTransformations(b *testing.B) {
	// Setup test matrix
	m := I4

	b.Run("Translate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultMat = m.Translate(2, 3, 4)
		}
	})

	b.Run("Scale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultMat = m.Scale(2, 3, 4)
		}
	})

	b.Run("RotateX", func(b *testing.B) {
		rad := math.Pi / 4
		for i := 0; i < b.N; i++ {
			resultMat = m.RotateX(rad)
		}
	})

	b.Run("RotateY", func(b *testing.B) {
		rad := math.Pi / 4
		for i := 0; i < b.N; i++ {
			resultMat = m.RotateY(rad)
		}
	})

	b.Run("RotateZ", func(b *testing.B) {
		rad := math.Pi / 4
		for i := 0; i < b.N; i++ {
			resultMat = m.RotateZ(rad)
		}
	})

	b.Run("Shear", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultMat = m.Shear(1, 2, 3, 4, 5, 6)
		}
	})
}

func BenchmarkTransformationChains(b *testing.B) {
	m := I4
	rad := math.Pi / 4

	b.Run("RotateScale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultMat = m.RotateZ(rad).Scale(2, 3, 4)
		}
	})

	b.Run("TranslateRotateScale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultMat = m.Translate(5, 6, 7).RotateZ(rad).Scale(2, 3, 4)
		}
	})

	b.Run("ComplexChain", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultMat = m.
				Translate(1, 2, 3).
				RotateX(rad).
				RotateY(rad).
				RotateZ(rad).
				Scale(2, 2, 2).
				Shear(1, 0, 0, 0, 0, 1)
		}
	})
}
