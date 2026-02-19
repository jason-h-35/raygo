package tracer

import "testing"

func BenchmarkColorCreation(b *testing.B) {
	b.Run("NewColorFromFloat64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = NewColorFromFloat64(0.8, 0.1, 0.3)
		}
	})
}

func BenchmarkColorOperations(b *testing.B) {
	c1 := LinearColor{0.8, 0.4, 0.2}
	c2 := LinearColor{0.2, 0.4, 0.8}
	f := float32(2)

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
	c := LinearColor{0.8, 0.4, 0.2}

	b.Run("RGBA", func(be *testing.B) {
		var r, g, b uint32
		for i := 0; i < be.N; i++ {
			r, g, b, _ = c.RGBA()
			results.scalar.float = float64(r + g + b)
		}
	})

	b.Run("ToPPMRange", func(b *testing.B) {
		var r, g, bl uint64
		for i := 0; i < b.N; i++ {
			r, g, bl = c.ToPPMRange(255)
			results.scalar.float = float64(r + g + bl)
		}
	})

	b.Run("AsFloats", func(be *testing.B) {
		var r, g, bl float32
		for i := 0; i < be.N; i++ {
			r, g, bl = c.AsFloats()
			results.scalar.float = float64(r + g + bl)
		}
	})

	b.Run("AsInts", func(be *testing.B) {
		var r, g, bl int
		for i := 0; i < be.N; i++ {
			r, g, bl = c.AsInts()
			results.scalar.float = float64(r + g + bl)
		}
	})
}

func BenchmarkColorComparisons(b *testing.B) {
	c1 := LinearColor{0.8, 0.4, 0.2}
	c2 := LinearColor{0.2, 0.4, 0.8}

	b.Run("Equals", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.scalar.bool = c1.Equals(c2)
		}
	})

	b.Run("Distance", func(b *testing.B) {
		var d float32
		for i := 0; i < b.N; i++ {
			d = c1.Distance(c2)
			results.scalar.float = float64(d)
		}
	})
}
