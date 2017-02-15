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
package fsm

import "log"
import "compiler/operation"
import "compiler/transformer"

type Transformer transformer.Transform
type VarTable map[string] int16
type TFTable map[string] Transformer

type FSM struct {
	tfstack *transformer.TFStack
	vartable *VarTable
	tftable *TFTable
}

type FSMError struct {
	oper string
	reason string
}

type VartableError struct {
	reason string
}

func NewVartableError(reason string) *VartableError {
	e := VartableError{reason}
	return &e
}

func NewFSMError(oper string, reason string) *FSMError {
	e := FSMError{oper,reason}
	return &e
}

func (e *FSMError) Error() string {
	return "FSM error: " + e.oper+ ": " + e.reason
}

func (e *VartableError) Error() string {
	return "Vartable error: " + e.reason
}

func NewVarTable() *VarTable {
	return &VarTable{}
}

func NewTFTable() *TFTable {
	return new(TFTable)
}

func NewFSM() *FSM {
	fsm := new(FSM)
	fsm.tfstack = transformer.NewTFStack()
	fsm.vartable = NewVarTable()
	fsm.tftable = NewTFTable()
	return fsm
}

func (vartable *VarTable) Assign(name string, v operation.Value) error {
	if v.Type == operation.VARIABLE {
		value,ok := (*vartable)[v.Name]
		if !ok {
			return NewVartableError("Undefined variable: "+v.Name)
		}
		(*vartable)[name] = value
		return nil
	} else if v.Type == operation.INTEGER {
		(*vartable)[name] = v.Number
		return nil
	}
	return NewVartableError("Invalid value: "+v.ToString())
}

func (fsm *FSM) Lookup(name string) (int16,bool) {
	value,ok := (*fsm.vartable)[name]
	return value,ok
}

func (fsm *FSM) Update(oper operation.Operation) error {
	switch oper.Command {
	case operation.UNDEFINED:
		if operation.Verbose {
			log.Output(1,"Undefined operation")
		}
		return NewFSMError(oper.ToString(),"undefined operation")
	case operation.LINE:
		fallthrough
	case operation.RECT:
		fallthrough
	case operation.CIRCLE:
		fallthrough
	case operation.OVAL:
		fallthrough
	case operation.POLYGON:
		return nil
	case operation.SET:
		err := fsm.vartable.Assign(oper.Name,oper.Args[0])
		if err != nil {
			return NewFSMError(oper.ToString(),err.Error())
		}
	}
	return nil
}
