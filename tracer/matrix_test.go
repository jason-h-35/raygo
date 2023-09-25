package tracer

import (
	"testing"
)

func Test_NewMat2(t *testing.T) {
	m := NewMat2([]float64{-3, 5, 1, -2})
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
	m := NewMat3([]float64{-3, 5, 0, 1, -2, -7, 0, 1, 1})
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
	m := NewMat4([]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5})
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
	m1 := NewMat4([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	m2 := NewMat4([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	if !m1.Equals(m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat3Equals(t *testing.T) {
	m1 := NewMat3([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2 := NewMat3([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if !m1.Equals(m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat2Equals(t *testing.T) {
	m1 := NewMat2([]float64{1, 2, 3, 4})
	m2 := NewMat2([]float64{1, 2, 3, 4})
	if !m1.Equals(m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat4TimesMat4(t *testing.T) {
	a := NewMat4([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	b := NewMat4([]float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8})
	expect := NewMat4([]float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42})
	result := a.TimesMat4(&b)
	if !result.Equals(expect) {
		t.Errorf("%v * %v should equal %v but was %v instead", a, b, expect, result)
	}
	identity := NewMat4([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	aID := a.TimesMat4(&identity)
	if !aID.Equals(a) {
		t.Errorf("expected %v times Identity Mat would be %v, but was %v", a, a, aID)
	}
}

func Test_Mat4TimesTuple(t *testing.T) {
	a := NewMat4([]float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1})
	b := NewTuple(1, 2, 3, 1)
	expect := NewTuple(18, 24, 33, 1)
	result := a.TimesTuple(b)
	if !result.Equals(expect) {
		t.Errorf("expected %v TimesTuple %v would be %v, but was %v", a, b, expect, result)
	}
}

func Test_Mat4Transpose(t *testing.T) {
	a := NewMat4([]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8})
	expect := NewMat4([]float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8})
	result := a.Transpose()
	if !result.Equals(expect) {
		t.Errorf("expected %v Transpose would be %v, but was %v", a, expect, result)
	}
	id := NewMat4([]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	idT := id.Transpose()
	if !id.Equals(idT) {
		t.Errorf("expected %v Transpose would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat2Determinant(t *testing.T) {
	a := NewMat2([]float64{1, 5, -3, 2})
	expect := 17.0
	result := a.Determinant()
	if abs(expect-result) > eps {
		t.Errorf("expected %v Determinant would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3SubMat(t *testing.T) {
	a := NewMat3([]float64{1, 5, 0, -3, 2, 7, 0, 6, 3})
	result := a.SubMat(0, 2)
	expect := NewMat2([]float64{-3, 2, 0, 6})
	if !result.Equals(expect) {
		t.Errorf("expected %v SubMat would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3Minor(t *testing.T) {
	a := NewMat3([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	result := a.Minor(1, 0)
	expect := 25.0
	if abs(result-expect) > eps {
		t.Errorf("expected %v Minor would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3Cofactor(t *testing.T) {
	a := NewMat3([]float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	result := a.Cofactor(1, 0)
	expect := -25.0
	if abs(result-expect) > eps {
		t.Errorf("expected %v Cofactor would be %v, but was %v", a, expect, result)
	}
}

func Test_Mat3Determinant(t *testing.T) {
	a := NewMat3([]float64{1, 2, 6, -5, 8, -4, 2, 6, 4})
	table := []struct {
		mat    Mat3
		fnName string
		fn     func(Mat3, int, int) float64
		is     int
		js     int
		expect float64
	}{
		{a, "Cofactor", Mat3.Cofactor, 0, 0, 56},
		{a, "Cofactor", Mat3.Cofactor, 0, 1, 12},
		{a, "Cofactor", Mat3.Cofactor, 0, 2, -46},
	}
	for _, it := range table {
		// calling fn on a mat in such a way the params can be extracted for the error msg
		result := it.fn(it.mat, it.is, it.js)
		if abs(result-it.expect) > 0 {
			t.Errorf("expected %v of %v at (%v, %v) would be %v but was %v",
				it.fnName, it.mat, it.is, it.js, it.expect, result)
		}
	}
	result := a.Determinant()
	expect := -196.0
	if abs(result-expect) > eps {
		t.Errorf("expected Determinant of %v would be %v but was %v", a, expect, result)
	}
}

func Test_Mat4Determinant(t *testing.T) {
	a := NewMat4([]float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9})
	table := []struct {
		mat    Mat4
		fnName string
		fn     func(Mat4, int, int) float64
		is     int
		js     int
		expect float64
	}{
		{a, "Cofactor", Mat4.Cofactor, 0, 0, 690},
		{a, "Cofactor", Mat4.Cofactor, 0, 1, 447},
		{a, "Cofactor", Mat4.Cofactor, 0, 2, 210},
		{a, "Cofactor", Mat4.Cofactor, 0, 3, 51},
	}
	for _, it := range table {
		// calling fn on a mat in such a way the params can be extracted for the error msg
		result := it.fn(it.mat, it.is, it.js)
		if abs(result-it.expect) > 0 {
			t.Errorf("expected %v of %v at (%v, %v) would be %v but was %v",
				it.fnName, it.mat, it.is, it.js, it.expect, result)
		}
	}
	result := a.Determinant()
	expect := -4071.0
	if abs(result-expect) > eps {
		t.Errorf("expected Determinant of %v would be %v but was %v", a, expect, result)
	}
}

func Test_Mat4CanInverse(t *testing.T) {
	a := NewMat4([]float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6})
	b := NewMat4([]float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0})
	if !a.CanInverse() {
		t.Errorf("%v should be invertible", a)
	}
	if b.CanInverse() {
		t.Errorf("%v should be invertible", b)
	}
}
