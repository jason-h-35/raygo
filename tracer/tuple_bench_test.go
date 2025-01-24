package tracer

import (
	"testing"
)

// Prevent compiler optimizations
var result Tuple

func BenchmarkTupleCreation(b *testing.B) {
	b.Run("NewPoint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = NewPoint(1.0, 2.0, 3.0)
		}
	})

	b.Run("NewVector", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = NewVector(1.0, 2.0, 3.0)
		}
	})
}

func BenchmarkTupleOperations(b *testing.B) {
	p1 := NewPoint(1.0, 2.0, 3.0)
	p2 := NewPoint(4.0, 5.0, 6.0)
	v1 := NewVector(1.0, 2.0, 3.0)

	b.Run("Plus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = p1.Plus(v1)
		}
	})

	b.Run("Minus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = p1.Minus(p2)
		}
	})

	b.Run("Times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = v1.Times(2.5)
		}
	})
}

func BenchmarkTupleVectorMath(b *testing.B) {
	v1 := NewVector(1.0, 2.0, 3.0)
	v2 := NewVector(4.0, 5.0, 6.0)

	b.Run("Length", func(b *testing.B) {
		var l float64
		for i := 0; i < b.N; i++ {
			l = v1.Length()
		}
		// Prevent optimization
		result = NewVector(l, 0, 0)
	})

	b.Run("Normalized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = v1.Normalized()
		}
	})

	b.Run("Dot", func(b *testing.B) {
		var d float64
		for i := 0; i < b.N; i++ {
			d = v1.Dot(v2)
		}
		// Prevent optimization
		result = NewVector(d, 0, 0)
	})

	b.Run("Cross", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = v1.Cross(v2)
		}
	})
}
