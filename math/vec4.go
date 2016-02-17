package math

// 4-dimensional vector.
type Vec4 [4]float32

func (vec *Vec4) X() float32 {
	return vec[0]
}

func (vec *Vec4) Y() float32 {
	return vec[1]
}

func (vec *Vec4) Z() float32 {
	return vec[2]
}

func (vec *Vec4) W() float32 {
	return vec[3]
}

func Vec4Zero() Vec4 {
	return Vec4{0.0, 0.0, 0.0, 0.0}
}
