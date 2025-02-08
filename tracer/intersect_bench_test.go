package tracer

import (
	"testing"
)

// Prevent compiler optimizations
var (
	resultRay  Ray
	resultInts []Intersect
	resultInt  Intersect
	resultBool bool
	resultSph  Sphere
)

func BenchmarkRayCreation(b *testing.B) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)

	b.Run("NewRay", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultRay = NewRay(origin, direction)
		}
	})

	b.Run("NewSphere", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultSph = NewSphere()
		}
	})
}

func BenchmarkRayOperations(b *testing.B) {
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	m := I4.Scale(2, 2, 2)

	b.Run("Position", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ray.Position(2.0)
		}
	})

	b.Run("Transform", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultRay = ray.Transform(m)
		}
	})
}

func BenchmarkIntersections(b *testing.B) {
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	sphere := NewSphere()
	xs := []Intersect{
		NewIntersect(sphere, 1),
		NewIntersect(sphere, 2),
		NewIntersect(sphere, -1),
	}

	b.Run("GetIntersects", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultInts = sphere.GetIntersects(ray)
		}
	})

	b.Run("Hit/Sorted", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultInt, resultBool = Hit(xs)
		}
	})

	b.Run("Hit/Unsorted", func(b *testing.B) {
		unsorted := []Intersect{
			NewIntersect(sphere, 2),
			NewIntersect(sphere, -1),
			NewIntersect(sphere, 1),
		}
		for i := 0; i < b.N; i++ {
			resultInt, resultBool = Hit(unsorted)
		}
	})
}

func BenchmarkIntersectComparisons(b *testing.B) {
	sphere := NewSphere()
	x1 := NewIntersect(sphere, 1)
	x2 := NewIntersect(sphere, 1)

	b.Run("Equals", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultBool = x1.Equals(x2)
		}
	})

	b.Run("Same", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultBool = x1.Same(x2)
		}
	})
}

func BenchmarkIntersectCreation(b *testing.B) {
	sphere := NewSphere()
	times := []float64{1, 2, 3, 4, 5}

	b.Run("NewIntersect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultInt = NewIntersect(sphere, 1.0)
		}
	})

	b.Run("NewIntersects", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultInts = NewIntersects(sphere, times...)
		}
	})
}
