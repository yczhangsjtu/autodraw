// This file is part of autodraw.
//
// Autodraw is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Autodraw is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with autodraw.  If not, see <http://www.gnu.org/licenses/>.
package transformer

import "testing"
import "math/rand"

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

func BenchmarkTransformCompose(b *testing.B) {
	tf1 := IdentityTransform()
	tf2 := NewTransform(1.5,0.0,0.0,0.0,1.5,0.0)
	tf3 := NewTransform(1.5,0.1,0.0,0.0,1.5,0.0)
	tf4 := NewTransform(2.25,0.15,0.0,0.0,2.25,0.0)
	tf5 := NewTransform(1.5,-0.1,0.0,0.0,1.5,0.0)
	tf6 := NewTransform(2.25,0.0,0.0,0.0,2.25,0.0)

	for i := 0; i < b.N; i++ {
		TransformCompose(tf1,tf2)
		TransformCompose(tf1,tf3)
		TransformCompose(tf1,tf4)
		TransformCompose(tf1,tf5)
		TransformCompose(tf1,tf6)
		TransformCompose(tf2,tf3)
		TransformCompose(tf2,tf4)
		TransformCompose(tf2,tf5)
		TransformCompose(tf2,tf6)
		TransformCompose(tf3,tf4)
		TransformCompose(tf3,tf5)
		TransformCompose(tf3,tf6)
		TransformCompose(tf4,tf5)
		TransformCompose(tf4,tf6)
		TransformCompose(tf5,tf6)
	}
}

func TestRotateTransform(t *testing.T) {
	ts := []float64 {
		0.1,0.2,0.3,0.4,0.5,0.6,0.7,0.8,0.9,1.0,1.5,2.0,2.5,3.0,3.14,
	}
	xs := []float64 {
		0.1,0.2,0.3,0.4,0.5,0.6,0.7,0.8,0.9,1.0,0.0,0.0,0.0,0.0,0.0,
	}
	ys := []float64 {
		0.4,0.3,0.2,0.1,0.0,-0.9,-0.8,-0.7,-0.6,1.0,0.0,1.0,1.0,1.0,1.0,
	}
	txs := []float64 {
		0.06,0.14,0.23,0.33,0.44,1.0,1.04,1.06,1.03,-0.3,0.0,-0.91,-0.6,-0.14,0.0,
	}
	tys := []float64 {
		0.41,0.33,0.28,0.25,0.24,-0.41,-0.16,0.09,0.33,1.38,0.0,-0.42,-0.8,-0.99,-1.0,
	}
	for i := 0; i < len(ts); i++ {
		tx,ty := ApplyTransform(RotateTransform(ts[i]),xs[i],ys[i])
		if txs[i] != tx || tys[i] != ty {
			t.Errorf("Rotate %f on [%f,%f], expect [%f,%f], got [%f,%f]",
				ts[i],xs[i],ys[i],txs[i],tys[i],tx,ty)
		}
	}
}

func TestTranslateTransform(t *testing.T) {
	for i := 0; i < 100; i++ {
		x0 := i2f(int16(rand.Int31())/2)
		y0 := i2f(int16(rand.Int31())/2)
		dx := i2f(int16(rand.Int31())/2)
		dy := i2f(int16(rand.Int31())/2)
		x1,y1 := ApplyTransform(TranslateTransform(dx,dy),x0,y0)
		if x1 != round(x0+dx) || y1 != round(y0+dy) {
			t.Errorf("Translate [%f,%f] by [%f,%f], expect [%f,%f], got [%f,%f]",
				x0,y0,dx,dy,round(x0+dx),round(y0+dy),x1,y1)
		}
	}
}
