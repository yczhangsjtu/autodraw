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

import "compiler/operation"
import "compiler/transformer"

type VarTable map[string] operation.Value
type TFTable map[string] *transformer.Transform

type FSM struct {
	tfstack *transformer.TFStack
	vartable *VarTable

	tmptransform *transformer.Transform
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
	return &TFTable{}
}

func NewFSM() *FSM {
	fsm := new(FSM)
	fsm.tfstack = transformer.NewTFStack()
	fsm.vartable = NewVarTable()
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
	} else if v.Type == operation.INTEGER || v.Type == operation.TRANSFORMER{
		(*vartable)[name] = v
		return nil
	}
	return NewVartableError("Invalid value: "+v.ToString())
}

func (fsm *FSM) Lookup(name string) (operation.Value,bool) {
	value,ok := (*fsm.vartable)[name]
	return value,ok && value.Type == operation.INTEGER
}

func (fsm *FSM) LookupValues(args []operation.Value) ([]int16,bool) {
	result := make([]int16,len(args))
	for i,v := range args {
		if v.Type == operation.VARIABLE {
			value,ok := fsm.Lookup(v.Name)
			if !ok {
				return result,false
			}
			result[i] = value.Number
		} else if v.Type == operation.INTEGER {
			result[i] = v.Number
		} else {
			return result,false
		}
	}
	return result,true
}

func (fsm *FSM) Update(oper operation.Operation) error {
	switch oper.Command {
	case operation.UNDEFINED:
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
	case operation.USE:
		value,ok := fsm.Lookup(oper.Name)
		if ok {
			if value.Type == operation.TRANSFORMER {
				fsm.tmptransform = value.Transform
			} else {
				return NewFSMError(oper.ToString()," "+oper.Name+" is not transform ")
			}
		} else {
			return NewFSMError(oper.ToString(),"undefined transform "+oper.Name)
		}
	case operation.PUSH:
	case operation.POP:
	case operation.TRANSFORM:
		tfvalues,ok := fsm.LookupValues(oper.Args)
		if !ok {
			return NewFSMError(oper.ToString(),"invalid transform arguments")
		}
		fsm.vartable.Assign(oper.Name,operation.NewTransformValue(ArgsToTransform(tfvalues)))
	case operation.DRAW:
	case operation.IMPORT:
	case operation.BEGIN:
	case operation.END:
	}
	return nil
}

func ArgsToTransform(args []int16) *transformer.Transform {
	return transformer.NewTransform(
		float64(args[0])/100.0,float64(args[1])/100.0,float64(args[2]),
		float64(args[3])/100.0,float64(args[4])/100.0,float64(args[5]),
	)
}
