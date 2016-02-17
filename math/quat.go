package math

type Quat [4]float32

func (vec *Quat) w() float32 {
	return vec[0]
}

func (vec *Quat) x() float32 {
	return vec[1]
}

func (vec *Quat) y() float32 {
	return vec[2]
}

func (vec *Quat) z() float32 {
	return vec[3]
}
