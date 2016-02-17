package math

type Vec2 [4]float32

func (vec *Vec4) x() float32 {
	return vec[0]
}

func (vec *Vec4) y() float32 {
	return vec[1]
}

func (vec *Vec4) z() float32 {
	return vec[2]
}

func (vec *Vec4) w() float32 {
	return vec[3]
}
