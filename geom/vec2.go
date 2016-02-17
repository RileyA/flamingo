package geom

// 2-dimensional vector.
type Vec2 [2]float32

func (vec *Vec2) X() float32 {
	return vec[0]
}

func (vec *Vec2) Y() float32 {
	return vec[1]
}

func Vec2Zero() Vec2 {
	return Vec2{0.0, 0.0}
}
