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
	"reflect"
)

type Operation struct {
	Command int16
	Name    string
	Args    []Value
}

func NewOperation(op int16) Operation {
	operation := Operation{}
	operation.Command = op
	return operation
}

func (op *Operation) Print() {
	fmt.Printf("%s", GetName(op.Command))
	if op.Name != "" {
		fmt.Printf(" %s", op.Name)
	}
	if len(op.Args) > 0 {
		fmt.Printf(" ")
		ValuesPrint(op.Args)
	}
}

func (op *Operation) ToString() string {
	ret := fmt.Sprintf("%s", GetName(op.Command))
	if op.Name != "" {
		ret += fmt.Sprintf(" %s", op.Name)
	}
	if len(op.Args) > 0 {
		ret += " "
		ret += ValuesToString(op.Args)
	}
	return ret
}

func (op *Operation) Equal(op2 Operation) bool {
	return reflect.DeepEqual(*op, op2)
}
