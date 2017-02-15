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

type TFStack struct {
	stack []*Transform
}

func NewTFStack() TFStack {
	return TFStack{}
}

func IsIdentity(tfstack *TFStack) bool {
	return len(tfstack.stack) == 0
}

func GetTransform(tfstack *TFStack) *Transform {
	if len(tfstack.stack) == 0 {
		identity := IdentityTransform()
		return identity
	}
	return tfstack.stack[len(tfstack.stack)-1]
}

func PushTransform(tfstack *TFStack, tf *Transform) {
	if len(tfstack.stack) == 0 {
		tfstack.stack = append(tfstack.stack,tf)
	} else {
		tf = GetTransform(tfstack).Compose(tf)
		tfstack.stack = append(tfstack.stack,tf)
	}
}

func PopTransform(tfstack *TFStack) bool {
	if len(tfstack.stack) == 0 {
		return false
	}
	tfstack.stack = tfstack.stack[:len(tfstack.stack)-1]
	return true
}
