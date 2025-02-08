package tracer

import (
	"testing"
)

func BenchmarkColorCreation(b *testing.B) {
	b.Run("NewColorFromFloat64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = NewColorFromFloat64(0.8, 0.1, 0.3)
		}
	})
}

func BenchmarkColorOperations(b *testing.B) {
	c1 := HDRColor{0x8000, 0x4000, 0x2000}
	c2 := HDRColor{0x2000, 0x4000, 0x8000}
	f := uint64(2)

	b.Run("Plus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = c1.Plus(c2)
		}
	})

	b.Run("Minus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = c1.Minus(c2)
		}
	})

	b.Run("Times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = c1.Times(f)
		}
	})

	b.Run("Hadamard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = c1.Hadamard(c2)
		}
	})
}

func BenchmarkColorConversions(b *testing.B) {
	c := HDRColor{0x8000, 0x4000, 0x2000}

	b.Run("RGBA", func(be *testing.B) {
		var r, g, b uint32
		for i := 0; i < be.N; i++ {
			r, g, b, _ = c.RGBA()
			// Store in results to prevent optimization
			results.color.c = HDRColor{uint64(r), uint64(g), uint64(b)}
		}
	})

	b.Run("ToPPMRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = c.ToPPMRange(255)
		}
	})

	b.Run("AsFloats", func(be *testing.B) {
		var r, g, b float64
		for i := 0; i < be.N; i++ {
			r, g, b = c.AsFloats()
			results.scalar.float = r + g + b // Store to prevent optimization
		}
	})

	b.Run("AsInts", func(be *testing.B) {
		var r, g, b int
		for i := 0; i < be.N; i++ {
			r, g, b = c.AsInts()
			results.scalar.float = float64(r + g + b) // Store to prevent optimization
		}
	})
}

func BenchmarkColorComparisons(b *testing.B) {
	c1 := HDRColor{0x8000, 0x4000, 0x2000}
	c2 := HDRColor{0x2000, 0x4000, 0x8000}

	b.Run("Equals", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.bool = c1.Equals(c2)
		}
	})

	b.Run("Distance", func(b *testing.B) {
		var d uint64
		for i := 0; i < b.N; i++ {
			d = c1.Distance(c2)
			results.color.c.R = d // Store to prevent optimization
		}
	})
}
