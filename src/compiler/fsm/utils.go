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

/*
Package fsm implements a simple Finite State Machine which takes operations as
inputs and updates its state. When the input operations are finished, the
generated instructions can be dumped to byte string.
*/
package fsm

import (
	"math"
	"compiler/transformer"
)

// ArgsToTransform given an array of six integers returns a transformer
// The transformer is a matrix of type float64, so to represent a transformer
// with some accuracy, use 1 in integer to represent 0.01 in float
func ArgsToTransform(args []int16) *transformer.Transform {
	return transformer.NewTransform(
		float64(args[0])/100.0, float64(args[1])/100.0, float64(args[2]),
		float64(args[3])/100.0, float64(args[4])/100.0, float64(args[5]),
	)
}

func ArgToRotate(arg int16) *transformer.Transform {
	return transformer.RotateTransform(
		float64(arg)/180.0*math.Pi,
	)
}

func ArgsToScale(args []int16) *transformer.Transform {
	return transformer.ScaleTransform(
		float64(args[0])/100.0, float64(args[1])/100.0,
	)
}

func ArgsToTranslate(args []int16) *transformer.Transform {
	return transformer.TranslateTransform(
		float64(args[0]), float64(args[1]),
	)
}
