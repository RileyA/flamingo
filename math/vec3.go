package math

type Vec3 [3]float32

func (vec *Vec3) x() float32 {
	return vec[0]
}

func (vec *Vec3) y() float32 {
	return vec[1]
}

func (vec *Vec3) z() float32 {
	return vec[2]
}
