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
package instruction

import "fmt"
import "strconv"
import "reflect"
import "compiler/operation"

type Instruction struct {
	Command int16
	Args []int16
}

type InstructionError struct {
	reason string
}

func NewInstruction() Instruction {
	return Instruction{0,[]int16{}}
}

func NewInstructionError(reason string) *InstructionError {
	e := InstructionError{reason}
	return &e
}

func (e *InstructionError) Error() string {
	return e.reason
}

func (inst *Instruction) Equal(inst2 Instruction) bool {
	return reflect.DeepEqual(*inst,inst2)
}

func (inst *Instruction) ToString() string {
	return fmt.Sprintf("%s %d",operation.OperationNames[inst.Command],inst.Args)
}

func (inst *Instruction) ToBytes() []byte {
	ret := make([]byte,len(inst.Args)*2+2)
	ret[0] = byte(inst.Command/256)
	ret[1] = byte(inst.Command%256)
	for i := 0; i < len(inst.Args); i++ {
		ret[2*i+2] = byte(uint(inst.Args[i])/256)
		ret[2*i+3] = byte(uint(inst.Args[i])%256)
	}
	return ret
}

func GetInstruction(command int16, args []int16) (Instruction,error) {
	switch command {
	case operation.LINE:
		fallthrough
	case operation.RECT:
		fallthrough
	case operation.CIRCLE:
		fallthrough
	case operation.OVAL:
		if len(args) != int(operation.ExpectArgNum(command)) {
			return NewInstruction(),NewInstructionError(
				"invalid number of arguments: "+operation.OperationNames[command]+
			  " requires "+strconv.Itoa(int(operation.ExpectArgNum(command)))+
				" args, got "+strconv.Itoa(len(args)))
		}
		inst := NewInstruction()
		inst.Command = command
		inst.Args = args
		return inst,nil
	case operation.POLYGON:
		if len(args) < 4 || len(args)%2 == 1 {
		return NewInstruction(),NewInstructionError(
			"invalid number of arguments: got "+
			strconv.Itoa(len(args))+" for polygon")
		}
		inst := NewInstruction()
		inst.Command = command
		inst.Args = append([]int16{int16(len(args))},args...)
		return inst,nil
	default:
		return NewInstruction(),NewInstructionError(
			"invalid draw command: "+operation.OperationNames[command])
	}
}
