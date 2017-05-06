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
	"strconv"
	"compiler/instruction"
	"compiler/operation"
	"compiler/transformer"
)


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
// Methods for FSM class //////////////////////////////////////////////////////
func (fsm *FSM) appendOperation(oper operation.Operation) {
	(*fsm.opertable)[fsm.current] = append((*fsm.opertable)[fsm.current],oper)
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
