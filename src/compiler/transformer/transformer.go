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

import "fmt"
import "math"

var Tolerance float64 = 1e-10

type Transform struct {
	matrix [3][3]float64
}

func NewTransform(t11,t12,t13,t21,t22,t23 float64) *Transform {
	transform := Transform{}
	transform.matrix = [3][3]float64 {
		{t11,t12,t13},{t21,t22,t23},{0,0,1.0}}
	return &transform
}

func (tf *Transform) PrintTransform() {
	fmt.Printf("[[%f,%f,%f]\n[%f,%f,%f]\n[%f,%f,%f]]\n",
		tf.matrix[0][0],tf.matrix[0][1],tf.matrix[0][2],
		tf.matrix[1][0],tf.matrix[1][1],tf.matrix[1][2],
		tf.matrix[2][0],tf.matrix[2][1],tf.matrix[2][2])
}

func (tf *Transform) ToString() string {
	return fmt.Sprintf("[[%f,%f,%f]\n[%f,%f,%f]\n[%f,%f,%f]]\n",
		tf.matrix[0][0],tf.matrix[0][1],tf.matrix[0][2],
		tf.matrix[1][0],tf.matrix[1][1],tf.matrix[1][2],
		tf.matrix[2][0],tf.matrix[2][1],tf.matrix[2][2])
}

func (tf *Transform) Equal(tf2 *Transform) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(tf.matrix[i][j]-tf2.matrix[i][j]) > Tolerance {
				return false
			}
		}
	}
	return true
}

func IdentityTransform() *Transform {
	return NewTransform(1.0,0.0,0.0,0.0,1.0,0.0)
}

func RotateTransform(t float64) *Transform {
	cost,sint := math.Cos(t),math.Sin(t)
	return NewTransform(cost,-sint,0,sint,cost,0)
}

func TranslateTransform(x,y float64) *Transform {
	return NewTransform(1,0,x,0,1,y)
}

func ScaleTransform(x,y float64) *Transform {
	return NewTransform(x,0,0,0,y,0)
}

func FlipxTransform() *Transform {
	return NewTransform(-1,0,0,0,1,0)
}

func FlipyTransform() *Transform {
	return NewTransform(1,0,0,0,-1,0)
}

func (tf1 *Transform) Compose(tf2 *Transform) *Transform {
	tf := new(Transform)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var sum float64 = 0.0
			for k := 0; k < 3; k++ {
				sum += tf1.matrix[i][k]*tf2.matrix[k][j]
			}
			tf.matrix[i][j] = sum
		}
	}
	return tf
}

func (tf *Transform) Apply3(x,y,z float64) (tx,ty,tz float64) {
	tx = tf.matrix[0][0]*x + tf.matrix[0][1]*y + tf.matrix[0][2]*z
	ty = tf.matrix[1][0]*x + tf.matrix[1][1]*y + tf.matrix[1][2]*z
	tz = tf.matrix[2][0]*x + tf.matrix[2][1]*y + tf.matrix[2][2]*z
	return
}

func (tf *Transform) Apply(x, y float64) (tx,ty float64) {
	tx,ty,tz := tf.Apply3(x,y,1.0)
	tx /= tz
	ty /= tz
	return
}
