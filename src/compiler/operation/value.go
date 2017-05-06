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
package operation

import (
	"fmt"
	"compiler/transformer"
)

type Value struct {
	Type      int16
	Name      string
	Number    int16
	Transform *transformer.Transform
}

func NewNumberValue(x int16) Value {
	return Value{INTEGER, "", x, nil}
}

func NewNumberValues(args ...int16) []Value {
	ret := make([]Value, len(args))
	for i, v := range args {
		ret[i] = NewNumberValue(v)
	}
	return ret
}

func NewVariableValue(name string) Value {
	if ValidName(name) {
		return Value{VARIABLE, name, 0, nil}
	}
	return Value{NAN, "", 0, nil}
}

func NewTransformValue(tf *transformer.Transform) Value {
	return Value{TRANSFORMER, "", 0, tf}
}

func (v *Value) Print() {
	switch v.Type {
	case INTEGER:
		fmt.Printf("%d", v.Number)
	case TRANSFORMER:
		v.Transform.Print()
	case VARIABLE:
		fmt.Printf("%s", v.Name)
	default:
		fmt.Printf("undefined")
	}
}

func ValuesPrint(vs []Value) {
	fmt.Printf("[")
	for i, v := range vs {
		if i > 0 {
			fmt.Printf(",")
		}
		v.Print()
	}
	fmt.Printf("]")
}

func (v *Value) ToString() string {
	switch v.Type {
	case INTEGER:
		return fmt.Sprintf("%d", v.Number)
	case TRANSFORMER:
		return v.Transform.ToString()
	case VARIABLE:
		return fmt.Sprintf("%s", v.Name)
	}
	return fmt.Sprintf("undefined")
}

func ValuesToString(vs []Value) string {
	ret := "["
	for i, v := range vs {
		if i > 0 {
			ret += ","
		}
		ret += v.ToString()
	}
	ret += "]"
	return ret
}

func ValuesToInt(values []Value) ([]int16, bool) {
	ret := make([]int16, len(values))
	for i, v := range values {
		if v.Type != INTEGER {
			return ret, false
		}
		ret[i] = v.Number
	}
	return ret, true
}
