package tracer

import (
	"math"
	"testing"
)

func Test_NewMat2(t *testing.T) {
	m := NewMat[Size2]([]float64{-3, 5, 1, -2})
	table := []MatVal{
		{0, 0, -3}, {0, 1, 5}, {1, 0, 1}, {1, 1, -2},
	}
	for _, v := range table {
		result := m.At(v.i, v.j)
		expect := v.val
		if result != expect {
			t.Errorf("expected %v at %v, %v. got %v", v.val, v.i, v.j, result)
		}
	}
}

func Test_NewMat3(t *testing.T) {
	m := NewMat[Size3]([]float64{-3, 5, 0, 1, -2, -7, 0, 1, 1})
	table := []MatVal{{0, 0, -3}, {1, 1, -2}, {2, 2, 1}}
	for _, v := range table {
		result := m.At(v.i, v.j)
		expect := v.val
		if result != expect {
			t.Errorf("expected %v at %v, %v. got %v", v.val, v.i, v.j, result)
		}
	}

}

func Test_NewMat4(t *testing.T) {
	m := NewMat[Size4]([]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5})
	table := []MatVal{
		{0, 0, 1.0}, {0, 3, 4.0}, {1, 0, 5.5}, {1, 2, 7.5},
		{2, 2, 11.0}, {3, 0, 13.5}, {3, 2, 15.5},
	}
	for _, v := range table {
		result := m.At(v.i, v.j)
		expect := v.val
		if result != expect {
			t.Errorf("expected %v at %v, %v. got %v", v.val, v.i, v.j, result)
		}
	}
}

func Test_Mat4Equals(t *testing.T) {
	m1 := NewMat[Size4]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	m2 := NewMat[Size4]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	if !m1.Equals(m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat3Equals(t *testing.T) {
	m1 := NewMat[Size3]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2 := NewMat[Size3]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if !m1.Equals(m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat2Equals(t *testing.T) {
	m1 := NewMat[Size2]([]float64{1, 2, 3, 4})
	m2 := NewMat[Size2]([]float64{1, 2, 3, 4})
	if !m1.Equals(m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat4TimesMat4(t *testing.T) {
	a := NewMat[Size4]([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	b := NewMat[Size4]([]float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8})
	expect := NewMat[Size4]([]float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42})
	result := a.Times(b)
	if !result.Equals(expect) {
		t.Errorf("%v * %v should equal %v but was %v instead", a, b, expect, result)
	}
	identity := NewMat[Size4]([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	aID := a.Times(identity)
	if !aID.Equals(a) {
		t.Errorf("expected %v times Identity Mat would be %v, but was %v", a, a, aID)
	}
}

func Test_Mat4TimesTuple(t *testing.T) {
	a := NewMat[Size4]([]float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1})
	b := NewTuple(1, 2, 3, 1)
	expect := NewTuple(18, 24, 33, 1)
	result := a.TimesTuple(b)
	if !result.Equals(expect) {
		t.Errorf("expected %v TimesTuple %v would be %v, but was %v", a, b, expect, result)
	}
}

func Test_Mat4Transpose(t *testing.T) {
	a := NewMat[Size4]([]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8})
	expect := NewMat[Size4]([]float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8})
	result := a.Transpose()
	if !result.Equals(expect) {
		t.Errorf("expected %v Transpose would be %v, but was %v", a, expect, result)
	}
	id := NewMat[Size4]([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	idT := id.Transpose()
	if !id.Equals(idT) {
		t.Errorf("expected %v Transpose would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat2Determinant(t *testing.T) {
	a := NewMat[Size2]([]float64{1, 5, -3, 2})
	expect := 17.0
	result := a.Determinant()
	if math.Abs(expect-result) > epsilon {
		t.Errorf("expected %v Determinant would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3SubMat(t *testing.T) {
	a := NewMat[Size3]([]float64{1, 5, 0, -3, 2, 7, 0, 6, 3})
	result := SubMat[Size3, Size2](a, 0, 2)
	expect := NewMat[Size2]([]float64{-3, 2, 0, 6})
	if !result.Equals(expect) {
		t.Errorf("expected %v SubMat would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3Minor(t *testing.T) {
	a := NewMat[Size3]([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	result := Minor(a, 1, 0)
	expect := 25.0
	if math.Abs(result-expect) > epsilon {
		t.Errorf("expected %v Minor would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3Cofactor(t *testing.T) {
	a := NewMat[Size3]([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	result := Cofactor(a, 1, 0)
	expect := -25.0
	if math.Abs(result-expect) > epsilon {
		t.Errorf("expected %v Cofactor would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3Determinant(t *testing.T) {
	a := NewMat[Size3]([]float64{1, 2, 6, -5, 8, -4, 2, 6, 4})
	table := []struct {
		mat    Mat[Size3]
		fnName string
		fn     func(Mat[Size3], int, int) float64
		is     int
		js     int
		expect float64
	}{
		{a, "Cofactor", Cofactor[Size3], 0, 0, 56},
		{a, "Cofactor", Cofactor[Size3], 0, 1, 12},
		{a, "Cofactor", Cofactor[Size3], 0, 2, -46},
	}
	for _, it := range table {
		// calling fn on a mat in such a way the params can be extracted for the error msg
		result := it.fn(it.mat, it.is, it.js)
		if math.Abs(result-it.expect) > 0 {
			t.Errorf("expected %v of %v at (%v, %v) would be %v but was %v",
				it.fnName, it.mat, it.is, it.js, it.expect, result)
		}
	}
	result := a.Determinant()
	expect := -196.0
	if math.Abs(result-expect) > epsilon {
		t.Errorf("expected Determinant of %v would be %v but was %v", a, expect, result)
	}
}

func Test_Mat4Determinant(t *testing.T) {
	a := NewMat[Size4]([]float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9})
	table := []struct {
		mat    Mat[Size4]
		fnName string
		fn     func(Mat[Size4], int, int) float64
		is     int
		js     int
		expect float64
	}{
		{a, "Cofactor", Cofactor[Size4], 0, 0, 690},
		{a, "Cofactor", Cofactor[Size4], 0, 1, 447},
		{a, "Cofactor", Cofactor[Size4], 0, 2, 210},
		{a, "Cofactor", Cofactor[Size4], 0, 3, 51},
	}
	for _, it := range table {
		// calling fn on a mat in such a way the params can be extracted for the error msg
		result := it.fn(it.mat, it.is, it.js)
		if math.Abs(result-it.expect) > 0 {
			t.Errorf("expected %v of %v at (%v, %v) would be %v but was %v",
				it.fnName, it.mat, it.is, it.js, it.expect, result)
		}
	}
	result := a.Determinant()
	expect := -4071.0
	if math.Abs(result-expect) > epsilon {
		t.Errorf("expected Determinant of %v would be %v but was %v", a, expect, result)
	}
}

func Test_Mat4CanInverse(t *testing.T) {
	a := NewMat[Size4]([]float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6})
	b := NewMat[Size4]([]float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0})
	if !a.CanInverse() {
		t.Errorf("%v should be invertible", a)
	}
	if b.CanInverse() {
		t.Errorf("%v should be invertible", b)
	}
}

func Test_Mat4Inverse(t *testing.T) {
	mats := []Mat[Size4]{
		NewMat[Size4]([]float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4}),
		NewMat[Size4]([]float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4}),
		NewMat[Size4]([]float64{9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2}),
	}

	inverses := []Mat[Size4]{
		NewMat[Size4]([]float64{0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639}),
		NewMat[Size4]([]float64{-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308}),
		NewMat[Size4]([]float64{-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333, -0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333}),
	}
	for i := range mats {
		result := mats[i].Inverse()
		if !result.Equals(inverses[i]) {
			t.Errorf("%v Inverse should be %v but was %v", mats[i], inverses[i], result)
		}
	}
}

func Test_Mat4InverseIdent(t *testing.T) {
	a := NewMat[Size4]([]float64{3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1})
	b := NewMat[Size4]([]float64{8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5})
	c := a.Times(b)
	bI := b.Inverse()
	result := c.Times(bI)
	if !result.Equals(a) {
		t.Errorf("A * B * B' should be A but instead was %v", result)
	}
}
