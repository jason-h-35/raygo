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

func (m *Mat4) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m *Mat3) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m *Mat2) At(i, j int) float64 {
	return m.vals[i][j]
}

func (m1 *Mat4) Equals(m2 *Mat4) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (m1 *Mat3) Equals(m2 *Mat3) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (m1 *Mat2) Equals(m2 *Mat2) bool {
	for i := range m1.vals {
		for j := range m1.vals {
			if abs(m1.vals[i][j]-m2.vals[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func (a *Mat4) TimesMat4(b *Mat4) Mat4 {
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

func (a *Mat4) TimesTuple(b Tuple) Tuple {
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
