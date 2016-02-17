package math

import "testing"

func TestVec3getters(t *testing.T) {
	vector := Vec3{1.0, 2.0, 3.0}
	if vector.X() != vector[0] {
		t.Errorf("X() getter returned: %f expected %f", vector.X(), vector[0])
	}
	if vector.Y() != vector[1] {
		t.Errorf("Y() getter returned: %f expected %f", vector.Y(), vector[1])
	}
	if vector.Z() != vector[2] {
		t.Errorf("Z() getter returned: %f expected %f", vector.Z(), vector[2])
	}
}
