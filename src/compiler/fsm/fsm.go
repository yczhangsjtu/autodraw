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

import "fmt"
import "math"
import "strconv"
import "compiler/instruction"
import "compiler/operation"
import "compiler/transformer"

type VarTable map[string] operation.Value
type OperationTable map[string] []operation.Operation

///////////////////////////////////////////////////////////////////////////////
// Definition of the FSM class ////////////////////////////////////////////////
type FSM struct {
	tfstack  *transformer.TFStack
	vartable *VarTable
	opertable *OperationTable
	instlist []instruction.Instruction

	tmptransform *transformer.Transform
	current string
	beginLevel int

	Verbose bool
}
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Following are error classes used in this package ///////////////////////////
type FSMError struct {
	oper   string
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
	e := FSMError{oper, reason}
	return &e
}

func (e *FSMError) Error() string {
	return "FSM error: " + e.oper + ": " + e.reason
}

func (e *VartableError) Error() string {
	return "Vartable error: " + e.reason
}

func (e *ArgError) Error() string {
	return "Argument error: " + e.reason
}
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Initialization functions ///////////////////////////////////////////////////

// NewVarTable generates an empty variable lookup table
func NewVarTable() *VarTable {
	return &VarTable{}
}

func NewOperationList() []operation.Operation{
	return []operation.Operation{}
}

func NewOperationTable() *OperationTable{
	return &OperationTable{}
}

// NewFSM generates an empty Finite State Machine and initializes the
// components
func NewFSM() *FSM {
	fsm := new(FSM)
	fsm.tfstack = transformer.NewTFStack()
	fsm.vartable = NewVarTable()
	fsm.opertable = NewOperationTable()
	return fsm
}
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Methods for VarTable class /////////////////////////////////////////////////

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
	} else if v.Type == operation.INTEGER || v.Type == operation.TRANSFORMER {
		(*vartable)[name] = v
		return nil
	}
	return NewVartableError("Invalid value: " + v.ToString())
}
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Methods for FSM class //////////////////////////////////////////////////////
func (fsm *FSM) appendOperation(oper operation.Operation) {
	(*fsm.opertable)[fsm.current] = append((*fsm.opertable)[fsm.current],oper)
}

/* Update is the most crucial method of FSM: takes an operation and update its
own state.

For simple drawing operations like LINE, RECT, POLYGON, the FSM
maps the coordinates using the current transformation matrix in the stack
and generate an instruction.

For operations like SET, TRANSFORM the FSM modifies its variable lookup
table.

For operations like USE, PUSH and POP the FSM modifies its matrix stack.
*/
func (fsm *FSM) Update(oper operation.Operation) error {
	// If there has been a BEGIN not yet ENDed, i.e. in a subfigure, just try to
	// log the operation into the corresponding operation list of the figure name
	if fsm.current != "" {
		switch oper.Command {
		// One more level of begin, doesn't have to evaluate it (that's the job of
		// the subfigure), but have to count the number of BEGINs to know which END
		// is the final END
		case operation.BEGIN:
			fsm.beginLevel++
			fsm.appendOperation(oper)
			return nil
		// Decrease a level of begin, of there is more than one level
		// If only one level, end the subfigure
		case operation.END:
			fsm.beginLevel--
			if fsm.beginLevel > 0 {
				fsm.appendOperation(oper)
			} else if fsm.beginLevel == 0 {
				fsm.current = ""
				return nil
			} else {
				return NewFSMError(oper.ToString(),"too many end operation")
			}
		// Ordinary operations, simply append
		default:
			fsm.appendOperation(oper)
			return nil
		}
	}
	switch oper.Command {
	case operation.UNDEFINED:
		return NewFSMError(oper.ToString(), "undefined operation")
	case operation.LINE:
		fallthrough
	case operation.RECT:
		fallthrough
	case operation.OVAL:
		fallthrough
	case operation.POLYGON:
		values, err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid drawing arguments: "+err.Error())
		}
		hasTmpTransform := fsm.tmptransform != nil
		if hasTmpTransform {
			fsm.tfstack.PushTransform(fsm.tmptransform)
			fsm.tmptransform = nil
		}
		values, err = fsm.ApplyTransform(values, oper.Command)
		if hasTmpTransform {
			fsm.tfstack.PopTransform()
		}
		if err != nil {
			return NewFSMError(
				oper.ToString(), "error in applying transform: "+err.Error())
		}
		inst, err := instruction.GetInstruction(oper.Command, values)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "error in generating instruction: "+err.Error())
		}
		if fsm.Verbose {
			fmt.Println(inst.ToString())
		}
		fsm.instlist = append(fsm.instlist, inst)
	case operation.SET:
		err := fsm.vartable.Assign(oper.Name, oper.Args[0])
		if err != nil {
			return NewFSMError(oper.ToString(), err.Error())
		}
	case operation.USE:
		value, ok := fsm.Lookup(oper.Name)
		if !ok {
			return NewFSMError(
				oper.ToString(), "failed to lookup transform "+oper.Name)
		}
		if value.Type != operation.TRANSFORMER {
			return NewFSMError(oper.ToString(), oper.Name+" is not transform")
		}
		fsm.tmptransform = value.Transform
	case operation.PUSH:
		value, ok := fsm.Lookup(oper.Name)
		if !ok {
			return NewFSMError(
				oper.ToString(), "failed to lookup transform "+oper.Name)
		}
		if value.Type != operation.TRANSFORMER {
			return NewFSMError(oper.ToString(), oper.Name+" is not transform")
		}
		fsm.tfstack.PushTransform(value.Transform)
	case operation.POP:
		ok := fsm.tfstack.PopTransform()
		if !ok {
			return NewFSMError(oper.ToString(), "stack already empty")
		}
	case operation.TRANSFORM:
		tfvalues, err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid rotate arguments: "+err.Error())
		}
		fsm.vartable.Assign(
			oper.Name, operation.NewTransformValue(ArgsToTransform(tfvalues)))
	case operation.ROTATE:
		tfvalues, err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid transform arguments: "+err.Error())
		}
		fsm.vartable.Assign(
			oper.Name, operation.NewTransformValue(ArgToRotate(tfvalues[0])))
	case operation.SCALE:
		tfvalues, err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid scale arguments: "+err.Error())
		}
		fsm.vartable.Assign(
			oper.Name, operation.NewTransformValue(ArgsToScale(tfvalues)))
	case operation.TRANSLATE:
		tfvalues, err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid scale arguments: "+err.Error())
		}
		fsm.vartable.Assign(
			oper.Name, operation.NewTransformValue(ArgsToTranslate(tfvalues)))
	case operation.DRAW:
		operlist,ok := (*fsm.opertable)[oper.Name]
		if !ok {
			return NewFSMError(
				oper.ToString(), "figure does not exist: "+oper.Name)
		}
		subfsm := NewFSM()
		subfsm.opertable = fsm.opertable
		hasTmpTransform := fsm.tmptransform != nil
		if hasTmpTransform {
			fsm.tfstack.PushTransform(fsm.tmptransform)
			fsm.tmptransform = nil
		}
		subfsm.PushTransform(fsm.tfstack.GetTransform())
		if hasTmpTransform {
			fsm.tfstack.PopTransform()
		}
		for _,suboper := range operlist {
			if fsm.Verbose {
				fmt.Printf("Subfigure %s: %s\n",oper.Name,suboper.ToString())
			}
			err := subfsm.Update(suboper)
			if err != nil {
				return NewFSMError(
					oper.ToString(),"error in figure "+oper.Name+":\n\t"+err.Error())
			}
		}
		fsm.instlist = append(fsm.instlist,subfsm.instlist...)
	case operation.IMPORT:
	case operation.BEGIN:
		_,ok := (*fsm.opertable)[oper.Name]
		if ok {
			return NewFSMError(
				oper.ToString(), "figure already exists: "+oper.Name)
		}
		(*fsm.opertable)[oper.Name] = NewOperationList()
		fsm.current = oper.Name
		fsm.beginLevel++
	case operation.END:
		return NewFSMError(oper.ToString(),"unexpected end of figure")
	}
	return nil
}

// FSM.Lookup is a wrapper around the lookup function of its variable table.
func (fsm *FSM) Lookup(name string) (operation.Value, bool) {
	value, ok := (*fsm.vartable)[name]
	return value, ok && (value.Type == operation.INTEGER ||
		value.Type == operation.TRANSFORMER)
}

// FSM.LookupValues takes an array of values which may contain unresolved
// variables and force them into an array of values of type INTEGER.
// Appearance of other types like TRANSFORMER will cause an error
func (fsm *FSM) LookupValues(args []operation.Value) ([]int16, error) {
	result := make([]int16, len(args))
	for i, v := range args {
		if v.Type == operation.VARIABLE {
			value, ok := fsm.Lookup(v.Name)
			if !ok {
				return result, NewVartableError("failed to lookup " + v.Name)
			}
			if value.Type != operation.INTEGER {
				return result, NewVartableError(v.Name + " is not integer")
			}
			result[i] = value.Number
		} else if v.Type == operation.INTEGER {
			result[i] = v.Number
		} else {
			return result, NewVartableError("invalid value type")
		}
	}
	return result, nil
}

// FSM.ApplyTransform apply the current transformation matrix to the
// coordinates list. The behavior is different for different drawing types.
//
// For LINE, POLYGON, take each pair of integers as (x,y) coordinates
// and apply the transformation.
//
// For RECT and OVAL, some kind of expansion has to be applied to the arguments
// so that the number of arguments is enough to represent the graph after
// transformation.
//
// For RECT, finally there will be 4 points stored, i.e. the four vertices
//
// For OVAL, finally 8 points are stored in instruction, which is enough to
// construct an oval after affine transformation using bezier curve
// TODO: in fact, four points are enough for oval
func (fsm *FSM) ApplyTransform(coords []int16, command int16) ([]int16, error) {
	result := make([]int16, len(coords))
	copy(result, coords)
	switch command {
	case operation.RECT:
		if len(coords) != 4 {
			return result, NewArgError(
				"invalid number of coordinates: " + strconv.Itoa(len(coords)))
		}
		result = make([]int16, 8)
		x1, y1, x2, y2 := float64(coords[0]), float64(coords[1]),
											float64(coords[2]), float64(coords[3])
		fx,fy := fsm.tfstack.GetTransform().Apply(float64(x1), float64(y1))
		result[0], result[1] = int16(fx), int16(fy)
		fx,fy = fsm.tfstack.GetTransform().Apply(float64(x1), float64(y2))
		result[2], result[3] = int16(fx), int16(fy)
		fx,fy = fsm.tfstack.GetTransform().Apply(float64(x2), float64(y2))
		result[4], result[5] = int16(fx), int16(fy)
		fx,fy = fsm.tfstack.GetTransform().Apply(float64(x2), float64(y1))
		result[6], result[7] = int16(fx), int16(fy)
		return result,nil
	case operation.LINE:
		fallthrough
	case operation.POLYGON:
		if len(coords) == 0 || len(coords)%2 == 1 {
			return result, NewArgError(
				"invalid number of coordinates: " + strconv.Itoa(len(coords)))
		}
		for i := 0; i < len(coords)/2; i++ {
			ix, iy := 2*i, 2*i+1
			x, y := coords[ix], coords[iy]
			fx, fy := fsm.tfstack.GetTransform().Apply(float64(x), float64(y))
			result[ix], result[iy] = int16(fx), int16(fy)
		}
		return result, nil
	case operation.OVAL:
		x, y, a, b := float64(coords[0]), float64(coords[1]),
									float64(coords[2]), float64(coords[3])
		result = make([]int16,16)
		tx, ty := x+a, y
		fx, fy := fsm.tfstack.GetTransform().Apply(tx,ty)
		result[0], result[1] = int16(fx), int16(fy)
		tx, ty = x+a, y+b
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[2], result[3] = int16(fx), int16(fy)
		tx, ty = x, y+b
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[4], result[5] = int16(fx), int16(fy)
		tx, ty = x-a, y+b
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[6], result[7] = int16(fx), int16(fy)
		tx, ty = x-a, y
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[8], result[9] = int16(fx), int16(fy)
		tx, ty = x-a, y-b
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[10], result[11] = int16(fx), int16(fy)
		tx, ty = x, y-b
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[12], result[13] = int16(fx), int16(fy)
		tx, ty = x+a, y-b
		fx, fy = fsm.tfstack.GetTransform().Apply(tx,ty)
		result[14], result[15] = int16(fx), int16(fy)
		return result, nil
	default:
		return result, NewArgError(
			"invalid draw command: " + operation.GetName(command))
	}
}

func (fsm *FSM) PushTransform(tf *transformer.Transform) {
	fsm.tfstack.PushTransform(tf)
}

func (fsm *FSM) PopTransform() error {
	if !fsm.tfstack.PopTransform() {
		return NewFSMError("pop", "stack already empty")
	}
	return nil
}

func (fsm *FSM) Assign(name string, v operation.Value) error {
	return fsm.vartable.Assign(name,v)
}

func (fsm *FSM) DumpInstructions() []byte {
	return instruction.InstructionsToBytes(fsm.instlist)
}
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Other utility functions ////////////////////////////////////////////////////

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
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
