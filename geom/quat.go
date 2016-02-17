package geom

// Simple quaternion type.
type Quat [4]float32

func (vec *Quat) W() float32 {
	return vec[0]
}

func (vec *Quat) X() float32 {
	return vec[1]
}

func (vec *Quat) Y() float32 {
	return vec[2]
}

func (vec *Quat) Z() float32 {
	return vec[3]
}

func QuatZero() Quat {
	return Quat{0.0, 0.0, 0.0, 0.0}
}

func QuatIdentity() Quat {
	return Quat{1.0, 0.0, 0.0, 0.0}
}
