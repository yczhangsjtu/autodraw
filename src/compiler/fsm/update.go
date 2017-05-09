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
	"fmt"
	"compiler/instruction"
	"compiler/operation"
)

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
		values, err := fsm.LookupIntegerValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid drawing arguments: "+err.Error())
		}
		oper.Args = values
		inst, err := instruction.OperationToInstruction(oper)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "error in getting instruction from operation: "+err.Error())
		}
		inst.Args, err = fsm.ApplyTmpTransform(inst.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "error in applying transform: "+err.Error())
		}
		if fsm.Verbose {
			fmt.Println(inst.ToString())
		}
		fsm.instlist = append(fsm.instlist, inst)
	case operation.TEXT:
		fallthrough
	case operation.NODE:
		values,err := fsm.LookupValues(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid text arguments: "+err.Error())
		}
		oper.Args = values
		inst, err := instruction.OperationToInstruction(oper)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "error in getting instruction from operation: "+err.Error())
		}
		pos,err := fsm.ApplyTmpTransform(inst.Args[0:2])
		if err != nil {
			return NewFSMError(
				oper.ToString(), "error in applying transform: "+err.Error())
		}
		inst.Args[0],inst.Args[1] = pos[0],pos[1]
		if fsm.Verbose {
			fmt.Println(inst.ToString())
		}
		fsm.instlist = append(fsm.instlist, inst)
	case operation.SET:
		err := fsm.Assign(oper.Name, oper.Args[0])
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
		tfvalues, err := fsm.LookupInts(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid rotate arguments: "+err.Error())
		}
		fsm.Assign(
			oper.Name, operation.NewTransformValue(ArgsToTransform(tfvalues)))
	case operation.ROTATE:
		tfvalues, err := fsm.LookupInts(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid transform arguments: "+err.Error())
		}
		fsm.Assign(
			oper.Name, operation.NewTransformValue(ArgToRotate(tfvalues[0])))
	case operation.SCALE:
		tfvalues, err := fsm.LookupInts(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid scale arguments: "+err.Error())
		}
		fsm.Assign(
			oper.Name, operation.NewTransformValue(ArgsToScale(tfvalues)))
	case operation.TRANSLATE:
		tfvalues, err := fsm.LookupInts(oper.Args)
		if err != nil {
			return NewFSMError(
				oper.ToString(), "invalid scale arguments: "+err.Error())
		}
		fsm.Assign(
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
		(*fsm.opertable)[oper.Name] = []operation.Operation{}
		fsm.current = oper.Name
		fsm.beginLevel++
	case operation.END:
		return NewFSMError(oper.ToString(),"unexpected end of figure")
	}
	return nil
}
