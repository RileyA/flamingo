package geom

import "math"

// 3-dimensional vector.
type Vec3 [3]float32

func (vec *Vec3) X() float32 {
	return vec[0]
}

func (vec *Vec3) Y() float32 {
	return vec[1]
}

func (vec *Vec3) Z() float32 {
	return vec[2]
}

func Vec3Zero() Vec3 {
	return Vec3{0.0, 0.0, 0.0}
}

func (lhs *Vec3) dot(rhs *Vec3) float32 {
	return lhs[0]*rhs[0] + lhs[1]*rhs[1] + lhs[2]*rhs[2]
}

func (vec *Vec3) length() float32 {
	return float32(math.Sqrt(float64(vec.dot(vec))))
}
