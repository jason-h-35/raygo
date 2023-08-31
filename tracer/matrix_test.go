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
	if !m1.Equals(&m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat3Equals(t *testing.T) {
	m1 := NewMat3([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2 := NewMat3([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if !m1.Equals(&m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}

func Test_Mat2Equals(t *testing.T) {
	m1 := NewMat4([]float64{1, 2, 3, 4})
	m2 := NewMat4([]float64{1, 2, 3, 4})
	if !m1.Equals(&m2) {
		t.Errorf("%v should equal %v", m1, m2)
	}
}
