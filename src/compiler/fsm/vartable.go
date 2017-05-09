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
	"compiler/operation"
)

type VarTable map[string] operation.Value

func NewVarTable() *VarTable {
	return &VarTable{}
}

// VarTable.Assign maps a string to a value. If the value is also a variable,
// lookup the variable name and map the string to the result found.
// If failed to find the variable, return an error.
// 
// If carried out successfully, the string will point to a value of type
// INTEGER or TRANSFORMER in this table.
func (vartable *VarTable) Assign(name string, v operation.Value) error {
	if v.Type == operation.VARIABLE {
		value, ok := (*vartable)[v.Name]
		if !ok {
			return NewVartableError("Undefined variable: " + v.Name)
		}
		(*vartable)[name] = value
		return nil
	} else if v.Type == operation.INTEGER || v.Type == operation.TRANSFORMER ||
	    v.Type == operation.STRING {
		(*vartable)[name] = v
		return nil
	}
	return NewVartableError("Invalid value: " + v.ToString())
}
