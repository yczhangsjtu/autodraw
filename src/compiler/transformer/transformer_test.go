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
import "math"
import "math/rand"

func TestTransformCompose(t *testing.T) {
	tf1 := IdentityTransform()
	tf2 := NewTransform(1.5,0.0,0.0,0.0,1.5,0.0)
	tf3 := NewTransform(1.5,0.1,0.0,0.0,1.5,0.0)
	tf4 := NewTransform(2.25,0.15,0.0,0.0,2.25,0.0)
	tf5 := NewTransform(1.5,-0.1,0.0,0.0,1.5,0.0)
	tf6 := NewTransform(2.25,0.0,0.0,0.0,2.25,0.0)

	tf12 := tf1.Compose(tf2)
	if !tf12.Equal(tf2) {
		t.Errorf("Expect\n %s.%s=%s, got %s",tf1.ToString(),
			tf2.ToString(),tf2.ToString(),tf12.ToString())
	}
	tf13 := tf1.Compose(tf3)
	if !tf13.Equal(tf3) {
		t.Errorf("Expect\n %s.%s=%s, got %s",tf1.ToString(),
			tf3.ToString(),tf3.ToString(),tf13.ToString())
	}
	tf23 := tf2.Compose(tf3)
	if !tf23.Equal(tf4) {
		t.Errorf("Expect\n %s.%s=%s, got %s",tf2.ToString(),
			tf3.ToString(),tf4.ToString(),tf23.ToString())
	}
	tf35 := tf3.Compose(tf5)
	if !tf35.Equal(tf6) {
		t.Errorf("Expect\n %s.%s=%s, got %s",tf3.ToString(),
			tf5.ToString(),tf6.ToString(),tf35.ToString())
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
		tf1.Compose(tf2)
		tf1.Compose(tf3)
		tf1.Compose(tf4)
		tf1.Compose(tf5)
		tf1.Compose(tf6)
		tf2.Compose(tf3)
		tf2.Compose(tf4)
		tf2.Compose(tf5)
		tf2.Compose(tf6)
		tf3.Compose(tf4)
		tf3.Compose(tf5)
		tf3.Compose(tf6)
		tf4.Compose(tf5)
		tf4.Compose(tf6)
		tf5.Compose(tf6)
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
	for i,tt := range ts {
		tx,ty := RotateTransform(tt).Apply(xs[i],ys[i])
		txs := math.Cos(tt)*xs[i]-math.Sin(tt)*ys[i]
		tys := math.Sin(tt)*xs[i]+math.Cos(tt)*ys[i]
		if txs != tx || tys != ty {
			t.Errorf("Rotate %f on [%f,%f], expect [%f,%f], got [%f,%f]",
				tt,xs[i],ys[i],txs,tys,tx,ty)
		}
	}
}

func TestTranslateTransform(t *testing.T) {
	for i := 0; i < 100; i++ {
		x0 := rand.Float64()
		y0 := rand.Float64()
		dx := rand.Float64()
		dy := rand.Float64()
		x1,y1 := TranslateTransform(dx,dy).Apply(x0,y0)
		if x1 != x0+dx || y1 != y0+dy {
			t.Errorf("Translate [%f,%f] by [%f,%f], expect [%f,%f], got [%f,%f]",
				x0,y0,dx,dy,x0+dx,y0+dy,x1,y1)
		}
	}
}
