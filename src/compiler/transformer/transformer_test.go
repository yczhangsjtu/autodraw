package transformer

import "testing"

func TestTransformCompose(t *testing.T) {
	tf1 := IdentityTransform()
	tf2 := NewTransform(1.5,0.0,0.0,0.0,1.5,0.0)
	tf3 := NewTransform(1.5,0.1,0.0,0.0,1.5,0.0)
	tf4 := NewTransform(2.25,0.15,0.0,0.0,2.25,0.0)
	tf5 := NewTransform(1.5,-0.1,0.0,0.0,1.5,0.0)
	tf6 := NewTransform(2.25,0.0,0.0,0.0,2.25,0.0)

	tf12 := TransformCompose(tf1,tf2)
	if !TransformEqual(tf12,tf2) {
		t.Errorf("Expect\n %s.%s=%s, got %s",TransformToString(tf1),
			TransformToString(tf2),TransformToString(tf2),TransformToString(tf12))
	}
	tf13 := TransformCompose(tf1,tf3)
	if !TransformEqual(tf13,tf3) {
		t.Errorf("Expect\n %s.%s=%s, got %s",TransformToString(tf1),
			TransformToString(tf3),TransformToString(tf3),TransformToString(tf13))
	}
	tf23 := TransformCompose(tf2,tf3)
	if !TransformEqual(tf23,tf4) {
		t.Errorf("Expect\n %s.%s=%s, got %s",TransformToString(tf2),
			TransformToString(tf3),TransformToString(tf4),TransformToString(tf23))
	}
	tf35 := TransformCompose(tf3,tf5)
	if !TransformEqual(tf35,tf6) {
		t.Errorf("Expect\n %s.%s=%s, got %s",TransformToString(tf3),
			TransformToString(tf5),TransformToString(tf6),TransformToString(tf35))
	}
}

