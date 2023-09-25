package tracer

// https://github.com/golang/go/issues/44253

type Mat2 struct {
	vals [2][2]float64
}

type Mat3 struct {
	vals [3][3]float64
}

type Mat4 struct {
	vals [4][4]float64
}

type MatVal struct {
	i   int
	j   int
	val float64
}

func NewMatVal(i int, j int, val float64) MatVal {
	return MatVal{i, j, val}
}

func NewMat2(a []float64) Mat2 {
	if len(a) != 4 {
		panic("Mat2 needs 4 elements.")
	}
	return Mat2{vals: [2][2]float64{
		{a[0], a[1]}, {a[2], a[3]}}}
}

func NewMat3(a []float64) Mat3 {
	if len(a) != 9 {
		panic("Mat3 needs 9 elements.")
	}
	return Mat3{vals: [3][3]float64{
		{a[0], a[1], a[2]},
		{a[3], a[4], a[5]},
		{a[6], a[7], a[8]},
	}}
}

func NewMat4(a []float64) Mat4 {
	if len(a) != 16 {
		panic("Mat4 needs 16 elements.")
	}
	return Mat4{vals: [4][4]float64{
		{a[0], a[1], a[2], a[3]},
		{a[4], a[5], a[6], a[7]},
		{a[8], a[9], a[10], a[11]},
		{a[12], a[13], a[14], a[15]},
	}}
}

func (m Mat4) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m Mat3) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m Mat2) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m1 Mat4) Equals(m2 Mat4) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (m1 Mat3) Equals(m2 Mat3) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (m1 Mat2) Equals(m2 Mat2) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (a Mat4) TimesMat4(b *Mat4) Mat4 {
	result := make([]float64, 0)
	for i := 0; i != 4; i++ {
		for j := 0; j != 4; j++ {
			next := 0.0
			for k := 0; k != 4; k++ {
				next += a.vals[i][k] * b.vals[k][j]
			}
			result = append(result, next)
		}
	}
	return NewMat4(result)
}

func (a Mat4) TimesTuple(b Tuple) Tuple {
	result := make([]float64, 0)
	bArr := b.AsArray()
	for i := 0; i != 4; i++ {
		next := 0.0
		for k := 0; k != 4; k++ {
			next += a.vals[i][k] * bArr[k]
		}
		result = append(result, next)
	}
	return NewTuple(result[0], result[1], result[2], result[3])
}

func (a Mat4) Transpose() Mat4 {
	av := a.vals
	return NewMat4([]float64{
		av[0][0], av[1][0], av[2][0], av[3][0],
		av[0][1], av[1][1], av[2][1], av[3][1],
		av[0][2], av[1][2], av[2][2], av[3][2],
		av[0][3], av[1][3], av[2][3], av[3][3],
	},
	)
}

func (a Mat2) Determinant() float64 {
	return a.vals[0][0]*a.vals[1][1] - a.vals[0][1]*a.vals[1][0]
}

func (a Mat3) SubMat(is, js int) Mat2 {
	s := make([]float64, 0)
	for i, row := range a.vals {
		for j, val := range row {
			if i != is && j != js {
				s = append(s, val)
			}
		}
	}
	return NewMat2(s)
}

func (a Mat4) SubMat(is, js int) Mat3 {
	s := make([]float64, 0)
	for i, row := range a.vals {
		for j, val := range row {
			if i != is && j != js {
				s = append(s, val)
			}
		}
	}
	return NewMat3(s)
}

func (a Mat3) Minor(is, js int) float64 {
	return a.SubMat(is, js).Determinant()
}

func (a Mat4) Minor(is, js int) float64 {
	return a.SubMat(is, js).Determinant()
}

func (a Mat3) Cofactor(is, js int) float64 {
	if (is+js)%2 == 1 {
		return -a.Minor(is, js)
	}
	return a.Minor(is, js)
}

func (a Mat4) Cofactor(is, js int) float64 {
	if (is+js)%2 == 1 {
		return -a.Minor(is, js)
	}
	return a.Minor(is, js)
}

func (a Mat3) Determinant() float64 {
	det, i := 0.0, 0 // may need to extract i later?
	for j, val := range a.vals[i] {
		det += val * a.Cofactor(i, j)
	}
	return det
}

func (a Mat4) Determinant() float64 {
	det, i := 0.0, 0 // may need to extract i later?
	for j, val := range a.vals[i] {
		det += val * a.Cofactor(i, j)
	}
	return det
}

func (a Mat3) CanInverse() bool {
	if abs(a.Determinant()) > eps {
		return true
	}
	return false
}

func (a Mat4) CanInverse() bool {
	if abs(a.Determinant()) > eps {
		return true
	}
	return false
}

func (a Mat3) Inverse() Mat3 {
	if !a.CanInverse() {
		panic("can't inverse this matrix")
	}
	inverse := NewMat3(make([]float64, 9))
	for i, row := range a.vals {
		for j := range row {
			inverse.vals[j][i] = a.Cofactor(i, j) / a.Determinant()
		}
	}
	return inverse
}

func (a Mat4) Inverse() Mat4 {
	if !a.CanInverse() {
		panic("can't inverse this matrix")
	}
	inverse := NewMat4(make([]float64, 16))
	for i, row := range a.vals {
		for j := range row {
			inverse.vals[j][i] = a.Cofactor(i, j) / a.Determinant()
		}
	}
	return inverse
}
