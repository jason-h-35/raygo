package tracer

import (
	"image/color"
	"testing"
)

func BenchmarkCanvasCreation(b *testing.B) {
	b.Run("NewCanvas/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.canvas.c = NewCanvas(100, 100)
		}
	})

	b.Run("NewCanvas/Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.canvas.c = NewCanvas(500, 500)
		}
	})

	b.Run("NewCanvas/Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.canvas.c = NewCanvas(1000, 1000)
		}
	})
}

func BenchmarkCanvasOperations(b *testing.B) {
	canvas := NewCanvas(100, 100)
	clr := HDRColor{0x8000, 0x4000, 0x2000}
	rgbaClr := color.RGBA{128, 64, 32, 255}

	b.Run("At/InBounds", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = canvas.AtHDR(50, 50)
		}
	})

	b.Run("At/OutOfBounds", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.color.c = canvas.AtHDR(150, 150)
		}
	})

	b.Run("Set/HDRColor", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			canvas.SetColor(50, 50, clr)
		}
	})

	b.Run("Set/RGBAColor", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			canvas.Set(50, 50, rgbaClr)
		}
	})
}

func BenchmarkCanvasExport(b *testing.B) {
	canvas := NewCanvas(100, 100)
	// Fill canvas with some test data
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			canvas.SetColor(x, y, HDRColor{
				R: uint64(x * 256),
				G: uint64(y * 256),
				B: uint64((x + y) * 128),
			})
		}
	}

	b.Run("PPMString/255", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.canvas.str = canvas.PPMStr(255)
		}
	})

	b.Run("PPMString/65535", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			results.canvas.str = canvas.PPMStr(65535)
		}
	})
}
