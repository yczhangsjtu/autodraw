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

import (
  "fmt"
  "reflect"
)

type Instruction struct {
	Command int16
	Args []int16
}

func NewInstruction() *Instruction {
	return new(Instruction)
}

func (inst *Instruction) Equal(inst2 *Instruction) bool {
	return reflect.DeepEqual(*inst,*inst2)
}

func (inst *Instruction) ToString() string {
	return fmt.Sprintf("%s %d",GetName(inst.Command),inst.Args)
}

func (inst *Instruction) ToBytes() []byte {
	i := 0
	ret := make([]byte,len(inst.Args)*2+4)
	setUint16(ret,&i,uint16(inst.Command))
	setUint16(ret,&i,uint16(len(inst.Args)))
	for ; i < len(inst.Args)+2; {
		setUint16(ret,&i,uint16(inst.Args[i-2]))
	}
	return ret
}

func GetInstance(command int16, args []int16) (*Instruction,error) {
	inst := NewInstruction()
	inst.Command = command
	inst.Args = args
	if !ExpectArgNum(command,len(args)) {
		return inst,NewInstructionError("invalid number of arguments for "+GetName(command))
	}
	return inst,nil;
}
