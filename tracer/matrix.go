package tracer

type Mat4 struct {
	arr [4][4]float64
}

type Mat3 struct {
	arr [3][3]float64
}

type Mat2 struct {
	arr [2][2]float64
}

func NewMat2(a, b, c, d float64) Mat2 {
	return Mat2{arr: [2][2]float64{{a, b}, {c, d}}}
}

func NewMat3(a, b, c, d, e, f, g, h, i float64) Mat3 {
	return Mat3{arr: [3][3]float64{{a, b, c}, {d, e, f}, {g, h, i}}}
}

func NewMat4(aa, ab, ac, ad, ba, bb, bc, bd, ca, cb, cc, cd, da, db, dc, dd float64) Mat4 {
	return Mat4{arr: [4][4]float64{{aa, ab, ac, ad}, {ba, bb, bc, bd}, {ca, cb, cc, cd}, {da, db, dc, dd}}}
}

func (m *Mat4) At(i, j int) float64 {
	return m.data[i][j]
}

func (m *Mat3) At(i, j int) float64 {
	return m.data[i][j]
}

func (m *Mat2) At(i, j int) float64 {
	return m.data[i][j]
}
