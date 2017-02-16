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

import "fmt"
import "strconv"
import "compiler/instruction"
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

type ArgError struct {
	reason string
}

func NewVartableError(reason string) *VartableError {
	e := VartableError{reason}
	return &e
}

func NewArgError(reason string) *ArgError {
	e := ArgError{reason}
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

func (e *ArgError) Error() string {
	return "Argument error: " + e.reason
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
	} else if v.Type == operation.INTEGER || v.Type == operation.TRANSFORMER {
		(*vartable)[name] = v
		return nil
	}
	return NewVartableError("Invalid value: "+v.ToString())
}

func (fsm *FSM) Lookup(name string) (operation.Value,bool) {
	value,ok := (*fsm.vartable)[name]
	return value,ok && (value.Type == operation.INTEGER ||
		value.Type == operation.TRANSFORMER)
}

func (fsm *FSM) LookupValues(args []operation.Value) ([]int16,error) {
	result := make([]int16,len(args))
	for i,v := range args {
		if v.Type == operation.VARIABLE {
			value,ok := fsm.Lookup(v.Name)
			if !ok {
				return result,NewVartableError("failed to lookup "+v.Name)
			}
			if value.Type != operation.INTEGER {
				return result,NewVartableError(v.Name+" is not integer")
			}
			result[i] = value.Number
		} else if v.Type == operation.INTEGER {
			result[i] = v.Number
		} else {
			return result,NewVartableError("invalid value type")
		}
	}
	return result,nil
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
		values,err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(),"invalid drawing arguments: "+err.Error())
		}
		hasTmpTransform := fsm.tmptransform != nil
		if hasTmpTransform {
			fsm.tfstack.PushTransform(fsm.tmptransform)
			fsm.tmptransform = nil
		}
		values,err = fsm.ApplyTransform(values,oper.Command)
		if hasTmpTransform {
			fsm.tfstack.PopTransform()
		}
		if err != nil {
			return NewFSMError(
				oper.ToString(),"error in applying transform: "+err.Error())
		}
		inst,err := instruction.GetInstruction(oper.Command,values)
		if err != nil {
			return NewFSMError(
				oper.ToString(),"error in generating instruction: "+err.Error())
		}
		fmt.Println(inst.ToString())
	case operation.SET:
		err := fsm.vartable.Assign(oper.Name,oper.Args[0])
		if err != nil {
			return NewFSMError(oper.ToString(),err.Error())
		}
	case operation.USE:
		value,ok := fsm.Lookup(oper.Name)
		if !ok {
			return NewFSMError(
				oper.ToString(),"failed to lookup transform "+oper.Name)
		}
		if value.Type != operation.TRANSFORMER {
			return NewFSMError(oper.ToString(),oper.Name+" is not transform")
		}
		fsm.tmptransform = value.Transform
	case operation.PUSH:
		value,ok := fsm.Lookup(oper.Name)
		if !ok {
			return NewFSMError(
				oper.ToString(),"failed to lookup transform "+oper.Name)
		}
		if value.Type != operation.TRANSFORMER {
			return NewFSMError(oper.ToString(),oper.Name+" is not transform")
		}
		fsm.tfstack.PushTransform(value.Transform)
	case operation.POP:
		ok := fsm.tfstack.PopTransform()
		if !ok {
			return NewFSMError(oper.ToString(),"stack already empty")
		}
	case operation.TRANSFORM:
		tfvalues,err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(),"invalid transform arguments: "+err.Error())
		}
		fsm.vartable.Assign(
			oper.Name,operation.NewTransformValue(ArgsToTransform(tfvalues)))
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

func (fsm *FSM)ApplyTransform(coords []int16, command int16) ([]int16,error) {
	result := make([]int16,len(coords))
	copy(result,coords)
	switch command {
	case operation.LINE:
		fallthrough
	case operation.RECT:
		fallthrough
	case operation.POLYGON:
		if len(coords) == 0 || len(coords)%2 == 1 {
			return result,NewArgError(
				"invalid number of coordinates: "+strconv.Itoa(len(coords)))
		}
		for i := 0; i < len(coords)/2; i++ {
			ix,iy := 2*i,2*i+1
			x,y := coords[ix],coords[iy]
			fx,fy := fsm.tfstack.GetTransform().Apply(float64(x),float64(y))
			result[ix],result[iy] = int16(fx),int16(fy)
		}
		return result,nil
	case operation.CIRCLE:
		fallthrough
	case operation.OVAL:
		x,y := coords[0],coords[1]
		fx,fy := fsm.tfstack.GetTransform().Apply(float64(x),float64(y))
		result[0],result[1] = int16(fx),int16(fy)
		return result,nil
	default:
		return result,NewArgError(
			"invalid draw command: "+operation.OperationNames[command])
	}
}
