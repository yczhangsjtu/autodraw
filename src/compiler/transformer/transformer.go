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
import "reflect"

type Transform struct {
	matrix [3][3]int16
}

func f2i(f float64) int16 {
	if f > 0.0 {
		return int16(f*100.0+0.5)
	} else {
		return int16(f*100.0-0.5)
	}
}

func i2f(i int16) float64 {
	return float64(i)/100.0;
}

func round(i float64) float64 {
	return i2f(f2i(i))
}

func NewTransform(t11,t12,t13,t21,t22,t23 float64) Transform {
	transform := Transform{}
	transform.matrix = [3][3]int16 {
		{f2i(t11),f2i(t12),f2i(t13)},
		{f2i(t21),f2i(t22),f2i(t23)},
		{0,0,f2i(1.0)}}
	return transform
}

func PrintTransform(tf Transform) {
	fmt.Printf("[[%f,%f,%f]\n[%f,%f,%f]\n[%f,%f,%f]]\n",
		i2f(tf.matrix[0][0]),i2f(tf.matrix[0][1]),i2f(tf.matrix[0][2]),
		i2f(tf.matrix[1][0]),i2f(tf.matrix[1][1]),i2f(tf.matrix[1][2]),
		i2f(tf.matrix[2][0]),i2f(tf.matrix[2][1]),i2f(tf.matrix[2][2]))
}

func TransformToString(tf Transform) string {
	return fmt.Sprintf("[[%f,%f,%f]\n[%f,%f,%f]\n[%f,%f,%f]]\n",
		i2f(tf.matrix[0][0]),i2f(tf.matrix[0][1]),i2f(tf.matrix[0][2]),
		i2f(tf.matrix[1][0]),i2f(tf.matrix[1][1]),i2f(tf.matrix[1][2]),
		i2f(tf.matrix[2][0]),i2f(tf.matrix[2][1]),i2f(tf.matrix[2][2]))
}

func TransformEqual(tf1,tf2 Transform) bool {
	return reflect.DeepEqual(tf1,tf2)
}

func IdentityTransform() Transform {
	return NewTransform(1.0,0.0,0.0,0.0,1.0,0.0)
}

func RotateTransform(t float64) Transform {
	cost,sint := math.Cos(t),math.Sin(t)
	return NewTransform(cost,-sint,0,sint,cost,0)
}

func TranslateTransform(x,y float64) Transform {
	return NewTransform(1,0,x,0,1,y)
}

func ScaleTransform(x,y float64) Transform {
	return NewTransform(x,0,0,0,y,0)
}

func FlipxTransform() Transform {
	return NewTransform(-1,0,0,0,1,0)
}

func FlipyTransform() Transform {
	return NewTransform(1,0,0,0,-1,0)
}

func TransformCompose(tf1,tf2 Transform) Transform {
	tf := Transform{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var sum float64 = 0.0
			for k := 0; k < 3; k++ {
				sum += i2f(tf1.matrix[i][k])*i2f(tf2.matrix[k][j])
			}
			tf.matrix[i][j] = f2i(sum)
		}
	}
	return tf
}

func ApplyTransform3(tf Transform, x,y,z float64) (tx,ty,tz float64) {
	tx = i2f(tf.matrix[0][0])*x + i2f(tf.matrix[0][1])*y + i2f(tf.matrix[0][2])*z
	ty = i2f(tf.matrix[1][0])*x + i2f(tf.matrix[1][1])*y + i2f(tf.matrix[1][2])*z
	tz = i2f(tf.matrix[2][0])*x + i2f(tf.matrix[2][1])*y + i2f(tf.matrix[2][2])*z
	return
}

func ApplyTransform(tf Transform, x, y float64) (tx,ty float64) {
	tx,ty,tz := ApplyTransform3(tf,x,y,1.0)
	tx = round(tx/tz)
	ty = round(ty/tz)
	return
}
